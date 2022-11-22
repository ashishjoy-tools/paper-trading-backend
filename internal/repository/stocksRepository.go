package repository

import (
	"github.com/ashishkujoy/paper-trading-backend/internal/model"
)

type StockRepository interface {
	GetStocksSymbol() ([]string, error)
	GetStockTradeHistory(symbol string) ([]model.StockTradeDetail, error)
	SaveStockTradeHistory(symbol string, history []model.StockTradeDetail) error
	GetStocksGist() ([]model.StockGist, error)
	AddSymbol(symbol string) error
	SaveStockGist(gist model.StockGist) error
}
