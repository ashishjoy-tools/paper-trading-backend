package main

import (
	"github.com/ashishkujoy/paper-trading-backend/internal"
	"github.com/ashishkujoy/paper-trading-backend/routers"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"os"
	"time"
)

func main() {
	r := gin.Default()
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_CONNECTION_URL"),
	})
	repository := internal.NewStockRepositoryRedisImpl(redisClient)
	stockService := internal.NewStockService(&repository)
	fetcher := internal.NewStockPriceFetcher(
		os.Getenv("STOCK_APP_KEY"),
		"alpha-vantage.p.rapidapi.com",
		os.Getenv("STOCK_APP_SERVER_URL"),
	)
	tradeManager := internal.NewStockTradeHistoryManagementService(
		&repository,
		4,
		time.Second,
		fetcher,
	)
	adminService := internal.NewAdminService(&tradeManager, &stockService)

	routers.NewStockRoutes(r, stockService)
	routers.NewAdminRoutes(r, &adminService)

	r.Run() // listen and serve on 0.0.0.0:8080
}
