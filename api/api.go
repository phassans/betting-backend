package api

import (
	"betting-backend/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func InitAPIs(db *gorm.DB, r *gin.Engine) {
	r.GET("/bets", func(c *gin.Context) {
		if err := model.GetBets(c, db); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	})

	// Retrieve a bet by ID
	r.GET("/bets/:id", func(c *gin.Context) {
		if err := model.GetBetByID(c, db); err != nil {
			if gorm.IsRecordNotFoundError(err) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Bet not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
	})

	// Retrieve all events
	r.GET("/events", func(c *gin.Context) {
		if err := model.GetEvents(c, db); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	})

	// Retrieve an event by ID
	r.GET("/events/:id", func(c *gin.Context) {
		if err := model.GetEventByID(c, db); err != nil {
			if gorm.IsRecordNotFoundError(err) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
	})

	// Create a new bet
	r.POST("/bets", func(c *gin.Context) {
		var bet model.Bet
		if err := c.BindJSON(&bet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := model.PostBet(db, bet); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, bet)
	})

	// Update a bet by ID
	r.PUT("/bets/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedBet model.Bet
		if err := c.BindJSON(&updatedBet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		betID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bet ID"})
			return
		}
		if err := model.UpdateBetByID(db, uint(betID), &updatedBet); err != nil {
			if gorm.IsRecordNotFoundError(err) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Bet not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, updatedBet)
	})
}
