package internal

type StockTradeDetail struct {
	Date          string `json:"date"`
	Open          string `json:"open"`
	High          string `json:"high"`
	Low           string `json:"low"`
	Close         string `json:"close"`
	AdjustedClose string `json:"adjustedClose"`
	Volume        string `json:"volume"`
}
