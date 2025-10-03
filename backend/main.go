package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Position struct {
	Ticker   string  `json:"ticker"`
	Amount   int     `json:"amount"`
	AvgPrice float64 `json:"avg_price"`
}

type Portfolio struct {
	Balance   float64    `json:"balance"`
	FreeFunds float64    `json:"free_funds"`
	Positions []Position `json:"positions"`
}

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/portfolio", getPortfolio)
		api.GET("/market-data/:ticker", getMarketData)
		api.POST("/order", createOrder)
	}

	r.Run(":8080")
}

// --- Handlers ---

func getPortfolio(c *gin.Context) {
	pf := Portfolio{
		Balance:   100000.0,
		FreeFunds: 75000.0,
		Positions: []Position{
			{Ticker: "SBER", Amount: 100, AvgPrice: 230.5},
			{Ticker: "YNDX", Amount: 10, AvgPrice: 2300.0},
		},
	}
	c.JSON(http.StatusOK, pf)
}

func getMarketData(c *gin.Context) {
	ticker := c.Param("ticker")
	// Мок: свечи
	candles := []map[string]interface{}{
		{"time": "2025-10-01", "open": 100.0, "close": 105.0, "high": 106.0, "low": 99.0, "volume": 12000},
		{"time": "2025-10-02", "open": 105.0, "close": 102.0, "high": 107.0, "low": 101.0, "volume": 8000},
	}
	c.JSON(http.StatusOK, gin.H{
		"ticker":  ticker,
		"candles": candles,
	})
}

func createOrder(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ticker, _ := payload["ticker"].(string)
	qty, _ := payload["quantity"].(float64)
	confirm, _ := payload["confirm"].(bool)

	if !confirm {
		c.JSON(http.StatusOK, gin.H{
			"status": "needs_confirmation",
			"ticker": ticker,
			"qty":    qty,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "order_created",
		"ticker":   ticker,
		"qty":      qty,
		"order_id": "ORD12345",
	})
}
