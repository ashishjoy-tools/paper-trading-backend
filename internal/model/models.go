package model

import (
	"encoding/json"
	"errors"
)

type StockPriceResponse struct {
	Symbol       string                      `json:"symbol"`
	HistoricData map[string]StockTradeDetail `json:"historicData"`
}

func ParseStockPriceResponse(res []byte, symbol string) (StockPriceResponse, error) {
	stockPriceResponse := StockPriceResponse{}
	resAsMap := make(map[string]interface{}, 0)
	err := json.Unmarshal(res, &resAsMap)
	if err != nil {
		return stockPriceResponse, err
	}
	historicalData, ok := resAsMap["Time Series (Daily)"]

	if !ok {
		return stockPriceResponse, errors.New("missing time series data in response")
	}
	historicalDataAsMap, ok := historicalData.(map[string]interface{})

	if !ok {
		return StockPriceResponse{}, errors.New("incorrect data type for time series data")
	}

	stockTradeDetails := make(map[string]StockTradeDetail, len(historicalDataAsMap))

	for date, tradeDetail := range historicalDataAsMap {
		tradeDetailAsMap := tradeDetail.(map[string]interface{})
		stockTradeDetails[date] = StockTradeDetail{
			Date:          date,
			Open:          tradeDetailAsMap["1. open"].(string),
			High:          tradeDetailAsMap["2. high"].(string),
			Low:           tradeDetailAsMap["3. low"].(string),
			Close:         tradeDetailAsMap["4. close"].(string),
			AdjustedClose: tradeDetailAsMap["5. adjusted close"].(string),
			Volume:        tradeDetailAsMap["6. volume"].(string),
		}
	}
	return StockPriceResponse{
		Symbol:       symbol,
		HistoricData: stockTradeDetails,
	}, nil
}
