package internal

import "fmt"

type StockTradeDetailsNotPresent struct {
	StockSymbol string `json:"stockSymbol"`
}

func (s *StockTradeDetailsNotPresent) Error() string {
	return fmt.Sprintf("Trade details not present for symbol: %s", s.StockSymbol)
}
