package services

import (
	"errors"
	"trader-backend/models"
)

// Интерфейс клиента Finam
type FinamClient interface {
	GetPortfolio() (*models.Portfolio, error)
	GetMarketData(ticker string) (interface{}, error)
	CreateOrder(payload map[string]interface{}) (interface{}, error)
}

// Простая реализация (mock)
type finamClientImpl struct{}

func NewFinamClient() FinamClient {
	return &finamClientImpl{}
}

func (f *finamClientImpl) GetPortfolio() (*models.Portfolio, error) {
	pf := &models.Portfolio{
		Balance:   100000.0,
		FreeFunds: 75000.0,
		Positions: []models.Position{
			{Ticker: "SBER", Amount: 100, AvgPrice: 230.5},
			{Ticker: "YNDX", Amount: 10, AvgPrice: 2300.0},
		},
	}
	return pf, nil
}

func (f *finamClientImpl) GetMarketData(ticker string) (interface{}, error) {
	example := map[string]interface{}{
		"ticker": ticker,
		"candles": []map[string]interface{}{
			{"time": "2025-10-01", "open": 100.0, "close": 105.0, "high": 106.0, "low": 99.0, "volume": 12000},
			{"time": "2025-10-02", "open": 105.0, "close": 102.0, "high": 107.0, "low": 101.0, "volume": 8000},
		},
	}
	return example, nil
}

func (f *finamClientImpl) CreateOrder(payload map[string]interface{}) (interface{}, error) {
	ticker, ok := payload["ticker"].(string)
	if !ok {
		return nil, errors.New("invalid ticker")
	}
	qty, _ := payload["quantity"].(float64)
	confirm, _ := payload["confirm"].(bool)

	if !confirm {
		return map[string]interface{}{
			"status": "needs_confirmation",
			"ticker": ticker,
			"qty":    qty,
		}, nil
	}

	return map[string]interface{}{
		"status":   "order_created",
		"ticker":   ticker,
		"qty":      qty,
		"order_id": "ORD12345",
	}, nil
}
