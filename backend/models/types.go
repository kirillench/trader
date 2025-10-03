package models

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
