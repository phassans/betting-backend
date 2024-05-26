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

// ReadEvents Read events from Ethereum and store in database
func ReadEvents(db *gorm.DB) {
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	contractAddress := common.HexToAddress(os.Getenv("BETTING_CONTRACT_ADDRESS"))
	bettingContract, err := betting.NewBetting(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Betting contract: %v", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Error: %v", err)
		case vLog := <-logs:
			handleLog(db, bettingContract, vLog)
		}
	}
}

// Handle logs and store in database
func handleLog(db *gorm.DB, contract *betting.Betting, vLog types.Log) {
	//eventName := contractABI.Events[vLog.Topics[0].Hex()].Name
	switch vLog.Topics[0].Hex() {
	case common.HexToHash(os.Getenv("TOPIC_BET_CREATED")).Hex(): // NewBetCreated event hash
		e, err := contract.ParseBetCreated(vLog)
		if err != nil {
			log.Printf("Failed to parse NewBetCreated event: %v", err)
			return
		}

		model.InsertBet(db, e)
		model.InsertEvent(e.BetID.Uint64(), "BetCreated", db, e.Raw)

	case common.HexToHash(os.Getenv("TOPIC_BET_ACTIVE")).Hex(): // BetActived event hash
		e, err := contract.ParseBetActive(vLog)
		if err != nil {
			log.Printf("Failed to parse BetActived event: %v", err)
			return
		}

		model.UpdateToActive(db, e) // Update bet status to active
		model.InsertEvent(e.BetID.Uint64(), "BetActive", db, e.Raw)

	case common.HexToHash(os.Getenv("TOPIC_BET_CLOSED")).Hex(): // BetClosed event hash
		e, err := contract.ParseBetClosed(vLog)
		if err != nil {
			log.Printf("Failed to parse BetClosed event: %v", err)
			return
		}

		model.UpdateToClosed(db, e)
		model.InsertEvent(e.BetID.Uint64(), "BetClosed", db, e.Raw)

	case common.HexToHash(os.Getenv("TOPIC_BET_REWARD_WITHDRAWAL")).Hex(): // RewardWithdrawal event hash
		e, err := contract.ParseBetRewardWithdrawal(vLog)
		if err != nil {
			log.Printf("Failed to parse RewardWithdrawal event: %v", err)
			return
		}

		model.UpdateToWithdraw(db, e)
		model.InsertEvent(e.BetID.Uint64(), "BetRewardWithdrawal", db, e.Raw)

	case common.HexToHash(os.Getenv("TOPIC_BET_REFUNDED")).Hex():
		e, err := contract.ParseBetRefunded(vLog)
		if err != nil {
			log.Printf("Failed to parse BetRefunded event: %v", err)
			return
		}
		model.InsertEvent(e.BetId.Uint64(), "BetRefunded", db, e.Raw)
	}
}
