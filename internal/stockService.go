package internal

type StockService struct {
	repository StockRepository
}

func (s *StockService) GetBySymbol(symbol string) {
	s.repository.GetStocksGist()
}

func (s *StockService) GetStocksGist() ([]StockGist, error) {
	return s.repository.GetStocksGist()
}

func (s *StockService) AddNewSymbol(symbol string) error {
	return s.repository.AddSymbol(symbol)
}

func NewStockService(repository StockRepository) StockService {
	return StockService{repository: repository}
}
