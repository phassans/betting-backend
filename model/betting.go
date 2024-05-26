package model

import (
	"betting-backend/betting"
	_ "betting-backend/betting"
	_ "github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

func InsertBet(db *gorm.DB, e *betting.BettingNewBetCreated) {
	db.Create(
		&Bet{
			BetID:       e.BetID.Uint64(),
			UserA:       e.UserA.Hex(),
			Amount:      floatFromBigInt(e.Amount),
			IsLong:      e.IsLong,
			CreateTime:  time.Unix(e.CreateTime.Int64(), 0),
			ExpireTime:  time.Unix(e.ExpireTime.Int64(), 0),
			ClosingTime: time.Unix(e.ClosingTime.Int64(), 0),
		})
}

func PostBet(db *gorm.DB, bet Bet) {
	db.Create(&bet)
}

func UpdateToActive(db *gorm.DB, e *betting.BettingBetActived) {
	db.Model(&Bet{}).Where("bet_id = ?", e.BetID.Uint64()).
		Updates(Bet{UserB: e.UserB.Hex(), OpeningPrice: floatFromBigInt(e.OpeningPrice), BetStatus: 1}) // Update bet status to active
}

func UpdateToClosed(db *gorm.DB, e *betting.BettingBetClosed) {
	db.Model(&Bet{}).Where("bet_id = ?", e.BetID.Uint64()).
		Updates(Bet{Winner: e.Winner.Hex(), Reward: floatFromBigInt(e.WinnerReward), ClosingPrice: floatFromBigInt(e.ClosingPrice), BetStatus: 2})

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
func GetBets(c *gin.Context, db *gorm.DB) {
	var bets []Bet
	if err := db.Find(&bets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bets)
}

// GetBetByID Get bet by ID
func GetBetByID(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var bet Bet
	if err := db.Where("bet_id = ?", id).First(&bet).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bet not found"})
		return
	}
	c.JSON(http.StatusOK, bet)
}

func floatFromBigInt(value *big.Int) float64 {
	f, _ := new(big.Float).SetInt(value).Float64()
	return f
}
