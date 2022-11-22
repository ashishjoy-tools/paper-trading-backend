package service

import (
	"github.com/ashishkujoy/paper-trading-backend/internal/model"
	"github.com/ashishkujoy/paper-trading-backend/internal/repository"
)

type StockService struct {
	repository repository.StockRepository
}

func (s *StockService) GetBySymbol(symbol string) {
	s.repository.GetStocksGist()
}

func (s *StockService) GetStocksGist() ([]model.StockGist, error) {
	return s.repository.GetStocksGist()
}

func (s *StockService) AddNewSymbol(symbol string) error {
	return s.repository.AddSymbol(symbol)
}

func NewStockService(repository repository.StockRepository) StockService {
	return StockService{repository: repository}
}
