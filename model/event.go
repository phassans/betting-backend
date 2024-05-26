package model

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type Event struct {
	ID          uint `gorm:"primary_key"`
	TxHash      string
	BlockNumber uint64
	BlockHash   string
	From        string
	InputData   []byte
	BetID       uint64
	EventType   string
	User        string
	Timestamp   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func InsertEvent(betID uint64, eventName string, db *gorm.DB, raw types.Log) {
	db.Create(
		&Event{TxHash: raw.TxHash.Hex(),
			BlockNumber: raw.BlockNumber,
			BlockHash:   raw.BlockHash.Hex(),
			User:        raw.Address.Hex(),
			InputData:   raw.Data,
			BetID:       betID,
			EventType:   eventName})
}

// GetEvents Get all events
func GetEvents(c *gin.Context, db *gorm.DB) {
	var events []Event
	if err := db.Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

// GetEventByID Get event by ID
func GetEventByID(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var event Event
	if err := db.Where("id = ?", id).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}
