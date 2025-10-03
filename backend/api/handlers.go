package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"trader-backend/services"
)

func HandleGetPortfolio(c *gin.Context, finam services.FinamClient) {
	pf, err := finam.GetPortfolio()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pf)
}

func HandleMarketData(c *gin.Context, finam services.FinamClient) {
	ticker := c.Param("ticker")
	data, err := finam.GetMarketData(ticker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func HandleCreateOrder(c *gin.Context, finam services.FinamClient) {
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := finam.CreateOrder(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
