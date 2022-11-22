package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ashishkujoy/paper-trading-backend/internal"
	"github.com/ashishkujoy/paper-trading-backend/internal/model"
	"github.com/go-redis/redis/v9"
	"time"
)

const RedisQueryTimeout = time.Second * 2

type StockRepositoryRedisImpl struct {
	redisClient *redis.Client
}

func (s *StockRepositoryRedisImpl) SaveStockGist(gist model.StockGist) error {
	ctx, cancelFunc := getContext()
	defer cancelFunc()
	bytes, err := json.Marshal(gist)
	if err != nil {
		return err
	}
	return s.redisClient.Set(ctx, fmt.Sprintf("STOCK_GIST_%s", gist.Symbol), bytes, redis.KeepTTL).Err()
}

func (s *StockRepositoryRedisImpl) AddSymbol(symbol string) error {
	ctx, cancelFunc := getContext()
	defer cancelFunc()
	stocksSymbol, err := s.GetStocksSymbol()
	if err != nil {
		return err
	}
	stocksSymbol = append(stocksSymbol, symbol)
	bytes, err := json.Marshal(stocksSymbol)
	if err != nil {
		return err
	}
	return s.redisClient.Set(ctx, "STOCK_SYMBOLS", bytes, redis.KeepTTL).Err()
}

func (s *StockRepositoryRedisImpl) GetStocksGist() ([]model.StockGist, error) {
	stockGists := make([]model.StockGist, 0)
	ctx, cancelFunc := getContext()
	defer cancelFunc()

	gistKeys, err := s.redisClient.Keys(ctx, "STOCK_GIST_*").Result()
	if err != nil {
		return nil, err
	}
	fmt.Printf("GIST KEYS %v\n", gistKeys)

	for _, gistKey := range gistKeys {
		ctxForGet, _ := getContext()
		bytes, err := s.redisClient.Get(ctxForGet, gistKey).Bytes()
		if err != nil {
			fmt.Printf("Error while fetching stock gist for %s, %v\n", gistKey, err)
			continue
		}
		var gist model.StockGist
		err = json.Unmarshal(bytes, &gist)
		if err != nil {
			fmt.Printf("Error while decoding stock gist %s, %v\n", gistKeys, err)
			continue
		}
		stockGists = append(stockGists, gist)
	}
	return stockGists, nil
}

func NewStockRepositoryRedisImpl(redisClient *redis.Client) StockRepositoryRedisImpl {
	return StockRepositoryRedisImpl{
		redisClient: redisClient,
	}
}

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), RedisQueryTimeout)
}

func (s *StockRepositoryRedisImpl) GetStocksSymbol() ([]string, error) {
	ctx, cancelFn := getContext()
	defer cancelFn()
	symbols := make([]string, 0)
	bytes, err := s.redisClient.Get(ctx, "STOCK_SYMBOLS").Bytes()
	if err != nil {
		if err == redis.Nil {
			ctx2, cancelFn2 := getContext()
			defer cancelFn2()

			s.redisClient.Set(
				ctx2,
				"STOCK_SYMBOLS",
				"[]",
				redis.KeepTTL,
			)
			return symbols, nil
		}
		return symbols, err
	}
	err = json.Unmarshal(bytes, &symbols)
	return symbols, err
}

func (s *StockRepositoryRedisImpl) GetStockTradeHistory(symbol string) ([]model.StockTradeDetail, error) {
	stockTradeDetails := make([]model.StockTradeDetail, 0)
	ctx, cancelFn := getContext()
	defer cancelFn()

	bytes, err := s.redisClient.Get(ctx, fmt.Sprintf("%s_STOCK_TRADE_DETAILS", symbol)).Bytes()

	if err != nil {
		if err == redis.Nil {
			return stockTradeDetails, &internal.StockTradeDetailsNotPresent{StockSymbol: symbol}
		}
		return stockTradeDetails, err
	}

	err = json.Unmarshal(bytes, &stockTradeDetails)

	return stockTradeDetails, err
}

func (s *StockRepositoryRedisImpl) SaveStockTradeHistory(symbol string, history []model.StockTradeDetail) error {
	ctx, cancelFunc := getContext()
	defer cancelFunc()

	bytes, err := json.Marshal(history)

	if err != nil {
		return err
	}

	return s.redisClient.Set(
		ctx,
		fmt.Sprintf("%s_STOCK_TRADE_DETAILS", symbol),
		bytes,
		redis.KeepTTL,
	).Err()
}
