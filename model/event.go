package model

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type Event struct {
	ID              uint `gorm:"primary_key"`
	TxHash          string
	BlockNumber     uint64
	BlockHash       string
	InputData       []byte
	BetID           uint64
	EventType       string
	ContractAddress string
	Timestamp       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func InsertEvent(betID uint64, eventName string, db *gorm.DB, raw types.Log) {
	db.Create(
		&Event{TxHash: raw.TxHash.Hex(),
			BlockNumber:     raw.BlockNumber,
			BlockHash:       raw.BlockHash.Hex(),
			ContractAddress: raw.Address.Hex(),
			InputData:       raw.Data,
			BetID:           betID,
			EventType:       eventName})
}

// GetEvents Get all events
func GetEvents(c *gin.Context, db *gorm.DB) error {
	var events []Event
	if err := db.Find(&events).Error; err != nil {
		return err
	}

	c.JSON(http.StatusOK, events)
	return nil
}

// GetEventByID Get event by ID
func GetEventByID(c *gin.Context, db *gorm.DB) error {
	id := c.Param("id")
	var event Event
	if err := db.Where("id = ?", id).First(&event).Error; err != nil {
		return err
	}

	c.JSON(http.StatusOK, event)
	return nil
}
