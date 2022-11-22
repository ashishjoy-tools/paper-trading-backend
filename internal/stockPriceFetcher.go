package internal

import (
	"io"
	"net/http"
	"strings"
)

type StockPriceFetcher struct {
	apiKey    string
	apiHost   string
	serverUrl string
}

func NewStockPriceFetcher(apiKey, apiHost, serverUrl string) StockPriceFetcher {
	return StockPriceFetcher{
		apiKey:    apiKey,
		apiHost:   apiHost,
		serverUrl: serverUrl,
	}
}

func (s *StockPriceFetcher) FetchDetailsForSymbol(symbol string) (StockPriceResponse, error) {
	var stockPriceResponse StockPriceResponse
	request, err := http.NewRequest(
		"GET",
		strings.Replace(s.serverUrl, "{SYMBOL}", symbol, 1),
		nil,
	)
	if err != nil {
		return stockPriceResponse, err
	}
	request.Header.Add("X-RapidAPI-Key", s.apiKey)
	request.Header.Add("X-RapidAPI-Host", "alpha-vantage.p.rapidapi.com")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return stockPriceResponse, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return stockPriceResponse, err
	}
	return ParseStockPriceResponse(body, symbol)
}
