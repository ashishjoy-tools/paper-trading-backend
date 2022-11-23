package model

type Stock struct {
	Symbol   string  `json:"symbol"`
	Quantity int     `json:"quantity"`
	BoughtAt float64 `json:"boughtAt"`
}

type StockWithCurrentValue struct {
	CurrentValue  float64 `json:"currentValue"`
	PerSharePrice float64 `json:"perSharePrice"`
	Stock
}

type Portfolio struct {
	Stocks []Stock
}

type PortfolioView struct {
	Username string                  `json:"username"`
	Balance  float64                 `json:"balance"`
	Stocks   []StockWithCurrentValue `json:"stocks"`
}

type Account struct {
	Username  string    `json:"username"`
	Balance   float64   `json:"balance"`
	Portfolio Portfolio `json:"portfolio"`
}
