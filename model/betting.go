package model

import (
	"betting-backend/betting"
	_ "betting-backend/betting"
	_ "github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"math/big"
	"net/http"
	"time"
)

type Bet struct {
	ID           uint   `gorm:"primary_key"`
	BetID        uint64 `gorm:"unique;not null"`
	UserA        string
	UserB        string
	Amount       float64
	Winner       string
	Reward       float64
	IsLong       bool
	CreateTime   time.Time
	ExpireTime   time.Time
	ClosingTime  time.Time
	OpeningPrice float64
	ClosingPrice float64
	BetStatus    int
}

// InsertBet inserts a new bet record into the database
func InsertBet(db *gorm.DB, e *betting.BettingBetCreated) error {
	bet := Bet{
		BetID:       e.BetID.Uint64(),
		UserA:       e.UserA.Hex(),
		Amount:      floatFromBigInt(e.Amount),
		IsLong:      e.IsLong,
		CreateTime:  time.Unix(e.CreateTime.Int64(), 0),
		ExpireTime:  time.Unix(e.ExpireTime.Int64(), 0),
		ClosingTime: time.Unix(e.ClosingTime.Int64(), 0),
	}
	if err := db.Create(&bet).Error; err != nil {
		log.Printf("Error inserting bet: %v", err)
		return err
	}

	return nil
}

// PostBet creates a new bet record in the database
func PostBet(db *gorm.DB, bet Bet) error {
	return db.Create(&bet).Error
}

// UpdateToActive updates a bet's status to active
func UpdateToActive(db *gorm.DB, e *betting.BettingBetActive) error {
	if err := db.Model(&Bet{}).Where("bet_id = ?", e.BetID.Uint64()).
		Updates(Bet{UserB: e.UserB.Hex(), OpeningPrice: floatFromBigInt(e.OpeningPrice), BetStatus: 1}).Error; err != nil {
		log.Printf("Error updating bet to active: %v", err)
		return err
	}

	return nil
}

// UpdateToClosed updates a bet's status to closed
func UpdateToClosed(db *gorm.DB, e *betting.BettingBetClosed) error {
	if err := db.Model(&Bet{}).Where("bet_id = ?", e.BetID.Uint64()).
		Updates(Bet{Winner: e.Winner.Hex(), Reward: floatFromBigInt(e.WinnerReward), ClosingPrice: floatFromBigInt(e.ClosingPrice), BetStatus: 2}).Error; err != nil {
		log.Printf("Error updating bet to closed: %v", err)
		return err
	}

	return nil
}

// UpdateToWithdraw updates a bet's status to withdrawn
func UpdateToWithdraw(db *gorm.DB, e *betting.BettingBetRewardWithdrawal) error {
	if err := db.Model(&Bet{}).Where("bet_id = ?", e.BetID.Uint64()).
		Updates(Bet{BetStatus: 3}).Error; err != nil {
		log.Printf("Error updating bet to withdrawn: %v", err)
		return err
	}

	return nil
}

// UpdateBetByID updates an existing bet in the database by ID
func UpdateBetByID(db *gorm.DB, id uint, updatedBet *Bet) error {
	var bet Bet

	if err := db.Where("bet_id = ?", id).First(&bet).Error; err != nil {
		return err
	}

	// Update fields
	bet.UserA = updatedBet.UserA
	bet.UserB = updatedBet.UserB
	bet.Amount = updatedBet.Amount
	bet.Winner = updatedBet.Winner
	bet.Reward = updatedBet.Reward
	bet.IsLong = updatedBet.IsLong
	bet.CreateTime = updatedBet.CreateTime
	bet.ExpireTime = updatedBet.ExpireTime
	bet.ClosingTime = updatedBet.ClosingTime
	bet.OpeningPrice = updatedBet.OpeningPrice
	bet.ClosingPrice = updatedBet.ClosingPrice
	bet.BetStatus = updatedBet.BetStatus

	return db.Save(&bet).Error
}

// GetBets Get all bets
func GetBets(c *gin.Context, db *gorm.DB) error {
	var bets []Bet
	if err := db.Find(&bets).Error; err != nil {
		return err
	}

	c.JSON(http.StatusOK, bets)
	return nil
}

// GetBetByID retrieves a bet by ID from the database
func GetBetByID(c *gin.Context, db *gorm.DB) error {
	id := c.Param("id")
	var bet Bet
	if err := db.Where("bet_id = ?", id).First(&bet).Error; err != nil {
		return err
	}

	c.JSON(http.StatusOK, bet)
	return nil
}

// floatFromBigInt converts a big.Int to a float64
func floatFromBigInt(value *big.Int) float64 {
	f, _ := new(big.Float).SetInt(value).Float64()
	return f
}
