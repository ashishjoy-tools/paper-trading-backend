package internal

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type StockTradeHistoryManagementService struct {
	stockRepository         StockRepository
	batchSize               int
	delayBetweenBatchUpdate time.Duration
	stockPriceFetcher       StockPriceFetcher
}

func NewStockTradeHistoryManagementService(
	repository StockRepository,
	batchSize int,
	delayBetweenBatchUpdate time.Duration,
	fetcher StockPriceFetcher,
) StockTradeHistoryManagementService {
	return StockTradeHistoryManagementService{
		stockRepository:         repository,
		batchSize:               batchSize,
		delayBetweenBatchUpdate: delayBetweenBatchUpdate,
		stockPriceFetcher:       fetcher,
	}
}

func (s *StockTradeHistoryManagementService) UpdateHistory() error {
	symbols, err := s.stockRepository.GetStocksSymbol()
	if err != nil {
		return err
	}
	symbolsBatched := toGroupOf(s.batchSize, symbols)
	for i, group := range symbolsBatched {
		fmt.Printf("Updating trade details for group: %d, %v\n", i, group)
		s.fetchStockDetailsFor(group)
		fmt.Printf("Done with updating details for group: %d, sleeping for %d\n", i, s.delayBetweenBatchUpdate)
		time.Sleep(s.delayBetweenBatchUpdate)
	}
	return nil
}

func (s *StockTradeHistoryManagementService) fetchStockDetailsFor(symbols []string) {
	wg := sync.WaitGroup{}

	for _, symbol := range symbols {
		wg.Add(1)
		go func(symbol string) {
			defer wg.Done()
			details, err := s.stockPriceFetcher.FetchDetailsForSymbol(symbol)
			if err != nil {
				fmt.Printf("An error occur while fetching details for symbol: %s\n%v\n", symbol, err)
			}
			history, err := s.stockRepository.GetStockTradeHistory(symbol)
			mergedDetails := mergeTradeData(details, history)
			err = s.stockRepository.SaveStockTradeHistory(symbol, mergedDetails)
			if err != nil {
				fmt.Printf("An error occur while saving details for symbol: %s\n%v\n", symbol, err)
			}
		}(symbol)
	}

	wg.Wait()
}

func mergeTradeData(details StockPriceResponse, history []StockTradeDetail) []StockTradeDetail {
	if len(history) == 0 {
		tradeDetails := make([]StockTradeDetail, len(details.HistoricData))
		i := 0
		for _, v := range details.HistoricData {
			tradeDetails[i] = v
			i++
		}
		return tradeDetails
	}
	lastTradeDate, _ := parseDate(history[len(history)-1].Date)

	for date, tradeDetail := range details.HistoricData {
		tradeDate, _ := parseDate(date)
		if lastTradeDate.Before(tradeDate) {
			history = append(history, tradeDetail)
		}
	}

	return history
}

func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

func toGroupOf(groupSize int, elements []string) [][]string {
	grouped := make([][]string, 0)

	for i := 0; i < len(elements); i += groupSize {
		endIndex := int(math.Min(float64(len(elements)), float64(i+groupSize)))
		grouped = append(grouped, elements[i:endIndex])
	}

	return grouped
}
