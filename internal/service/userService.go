package service

import (
	"github.com/ashishkujoy/paper-trading-backend/internal/model"
	"github.com/ashishkujoy/paper-trading-backend/internal/repository"
)

type UserService struct {
	userRepository  *repository.UserRepository
	stockRepository repository.StockRepository
}

func NewUserService(
	userRepository *repository.UserRepository,
	stockRepository repository.StockRepository,
) UserService {
	return UserService{userRepository: userRepository, stockRepository: stockRepository}
}

func (p *UserService) GetPortfolioForUser(username string) (model.PortfolioView, error) {
	portfolio := model.PortfolioView{}
	account, err := p.userRepository.GetUserAccount(username)
	if err != nil {
		return portfolio, err
	}

	portfolio.Username = account.Username
	portfolio.Balance = account.Balance

	stocks := make([]model.StockWithCurrentValue, len(account.Portfolio.Stocks))

	for index, stock := range account.Portfolio.Stocks {
		price, err := p.stockRepository.GetStockCurrentPrice(stock.Symbol)
		if err != nil {
			return portfolio, err
		}
		stocks[index] = model.StockWithCurrentValue{
			CurrentValue:  float64(stock.Quantity) * price,
			PerSharePrice: price,
			Stock:         stock,
		}
	}
	portfolio.Stocks = stocks
	return portfolio, nil
}

func (p *UserService) CreateNewUser(username string, openingBalance float64) (model.PortfolioView, error) {
	portfolio := model.PortfolioView{
		Username: username,
		Balance:  openingBalance,
		Stocks:   make([]model.StockWithCurrentValue, 0),
	}
	account := model.Account{
		Username:  username,
		Balance:   openingBalance,
		Portfolio: model.Portfolio{Stocks: make([]model.Stock, 0)},
	}
	err := p.userRepository.Save(account)
	return portfolio, err
}
