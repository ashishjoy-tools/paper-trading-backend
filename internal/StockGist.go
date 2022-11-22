package internal

import "strconv"

type StockGist struct {
	Symbol                          string  `json:"symbol"`
	Name                            string  `json:"name"`
	CurrentPrice                    float64 `json:"currentPrice"`
	ChangeFromPreviousDay           float64 `json:"changeFromPreviousDay"`
	PercentageChangeFromPreviousDay float64 `json:"percentageChangeFromPreviousDay"`
}

func StockGistFrom(tradeHistory []StockTradeDetail, symbol, name string) StockGist {
	stockGist := StockGist{Symbol: symbol, Name: name}

	if len(tradeHistory) > 1 {
		todayTrade := tradeHistory[len(tradeHistory)-1]
		previousDayTrade := tradeHistory[len(tradeHistory)-2]
		previousClose := toFloat(previousDayTrade.Close)
		todayClose := toFloat(todayTrade.Close)
		changeInClose := todayClose - previousClose

		stockGist.CurrentPrice = todayClose
		stockGist.ChangeFromPreviousDay = changeInClose
		stockGist.PercentageChangeFromPreviousDay = (changeInClose / previousClose) * 100
	} else {
		todayTrade := tradeHistory[len(tradeHistory)-1]

		stockGist.CurrentPrice = toFloat(todayTrade.Close)
		stockGist.ChangeFromPreviousDay = 0
		stockGist.PercentageChangeFromPreviousDay = 0
	}

	return stockGist
}

func toFloat(value string) float64 {
	parsed, _ := strconv.ParseFloat(value, 64)
	return parsed
}
