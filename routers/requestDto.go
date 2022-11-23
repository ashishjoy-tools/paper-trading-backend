package routers

type AddNewSymbolRequest struct {
	Symbol string `json:"symbol" binding:"required"`
}

type CreateNewUserRequest struct {
	Username string  `json:"username" binding:"required"`
	Balance  float64 `json:"balance" binding:"required"`
}

type StockBuySellRequest struct {
	Symbol   string `json:"symbol" binding:"required"`
	Action   string `json:"action" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}
