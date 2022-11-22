package internal

type StockRepository interface {
	GetStocksSymbol() ([]string, error)
	GetStockTradeHistory(symbol string) ([]StockTradeDetail, error)
	SaveStockTradeHistory(symbol string, history []StockTradeDetail) error
	GetStocksGist() ([]StockGist, error)
	AddSymbol(symbol string) error
}
