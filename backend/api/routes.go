package api

import (
	"github.com/gin-gonic/gin"
	"trader-backend/services"
)

func RegisterRoutes(r *gin.Engine, finam services.FinamClient) {
	api := r.Group("/api")
	{
		api.GET("/portfolio", func(c *gin.Context) { HandleGetPortfolio(c, finam) })
		api.GET("/market-data/:ticker", func(c *gin.Context) { HandleMarketData(c, finam) })
		api.POST("/order", func(c *gin.Context) { HandleCreateOrder(c, finam) })
	}
}
