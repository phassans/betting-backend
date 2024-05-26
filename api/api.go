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
		model.GetBets(c, db)
	})

	r.GET("/bets/:id", func(c *gin.Context) {
		model.GetBetByID(c, db)
	})

	r.GET("/events", func(c *gin.Context) {
		model.GetEvents(c, db)
	})

	r.GET("/events/:id", func(c *gin.Context) {
		model.GetEventByID(c, db)
	})

	// Create a new bet
	r.POST("/bets", func(c *gin.Context) {
		var bet model.Bet
		if err := c.BindJSON(&bet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		model.PostBet(db, bet)
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
