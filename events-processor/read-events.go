package events_processor

import (
	"betting-backend/betting"
	"betting-backend/model"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

// ReadEvents reads events from the Ethereum blockchain and stores them in the database.
func ReadEvents(db *gorm.DB) {
	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		log.Fatalf("RPC_URL is not set in environment variables")
	}

	// Connect to the Ethereum client using the RPC URL from environment variables.
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contractAddressStr := os.Getenv("BETTING_CONTRACT_ADDRESS")
	if contractAddressStr == "" {
		log.Fatalf("BETTING_CONTRACT_ADDRESS is not set in environment variables")
	}
	// Get the betting contract address from environment variables and instantiate the contract.
	contractAddress := common.HexToAddress(contractAddressStr)
	bettingContract, err := betting.NewBetting(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Betting contract: %v", err)
	}

	// Create a filter query to listen for events from the betting contract address.
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	// Create a channel to receive logs and subscribe to logs based on the filter query.
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}

	// Handle graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Listen for logs and handle them appropriately.
	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Error: %v", err)
		case vLog := <-logs:
			handleLog(db, bettingContract, vLog)
		case <-ctx.Done():
			log.Println("Shutting down event processor...")
			return
		}
	}
}

// handleLog processes the received log and stores the event in the database.
func handleLog(db *gorm.DB, contract *betting.Betting, vLog types.Log) {
	// 	Switch based on the event topic hash to identify the event type.
	switch vLog.Topics[0].Hex() {
	case common.HexToHash(os.Getenv("TOPIC_BET_CREATED")).Hex(): // NewBetCreated event hash
		e, err := contract.ParseBetCreated(vLog)
		if err != nil {
			log.Printf("Failed to parse NewBetCreated event: %v", err)
			return
		}

		// Insert the new bet and event into the database.
		model.InsertBet(db, e)
		model.InsertEvent(e.BetID.Uint64(), "BetCreated", db, e.Raw)

	case common.HexToHash(os.Getenv("TOPIC_BET_ACTIVE")).Hex(): // BetActived event hash
		e, err := contract.ParseBetActive(vLog)
		if err != nil {
			log.Printf("Failed to parse BetActived event: %v", err)
			return
		}

		// Update the bet status to active and insert the event into the database.
		model.UpdateToActive(db, e)
		model.InsertEvent(e.BetID.Uint64(), "BetActive", db, e.Raw)

	case common.HexToHash(os.Getenv("TOPIC_BET_CLOSED")).Hex(): // BetClosed event hash
		e, err := contract.ParseBetClosed(vLog)
		if err != nil {
			log.Printf("Failed to parse BetClosed event: %v", err)
			return
		}

		// Update the bet status to closed and insert the event into the database.
		model.UpdateToClosed(db, e)
		model.InsertEvent(e.BetID.Uint64(), "BetClosed", db, e.Raw)

	case common.HexToHash(os.Getenv("TOPIC_BET_REWARD_WITHDRAWAL")).Hex(): // RewardWithdrawal event hash
		e, err := contract.ParseBetRewardWithdrawal(vLog)
		if err != nil {
			log.Printf("Failed to parse RewardWithdrawal event: %v", err)
			return
		}

		// Update the bet status to withdraw and insert the event into the database.
		model.UpdateToWithdraw(db, e)
		model.InsertEvent(e.BetID.Uint64(), "BetRewardWithdrawal", db, e.Raw)

	case common.HexToHash(os.Getenv("TOPIC_BET_REFUNDED")).Hex():
		e, err := contract.ParseBetRefunded(vLog)
		if err != nil {
			log.Printf("Failed to parse BetRefunded event: %v", err)
			return
		}

		// insert the event into the database.
		model.InsertEvent(e.BetId.Uint64(), "BetRefunded", db, e.Raw)

	default:
		log.Printf("Unhandled event topic: %v", vLog.Topics[0].Hex())
	}
}
