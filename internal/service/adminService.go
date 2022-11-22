package service

type AdminService struct {
	tradeHistoryManager *StockTradeHistoryManagementService
	stockService        *StockService
}

func NewAdminService(
	tradeHistoryManager *StockTradeHistoryManagementService,
	stockService *StockService,
) AdminService {
	return AdminService{
		tradeHistoryManager: tradeHistoryManager,
		stockService:        stockService,
	}
}

func (s *AdminService) UpdateStockTradeDetails() error {
	return s.tradeHistoryManager.UpdateHistory()
}

func (s *AdminService) AddNewSymbol(symbol string) error {
	return s.stockService.AddNewSymbol(symbol)
}
