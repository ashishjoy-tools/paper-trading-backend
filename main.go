package main

import (
	"github.com/ashishkujoy/paper-trading-backend/internal/repository"
	"github.com/ashishkujoy/paper-trading-backend/internal/service"
	"github.com/ashishkujoy/paper-trading-backend/routers"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"os"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(routers.CORSMiddleware())
	
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_CONNECTION_URL"),
	})
	repository := repository.NewStockRepositoryRedisImpl(redisClient)
	stockService := service.NewStockService(&repository)
	fetcher := service.NewStockPriceFetcher(
		os.Getenv("STOCK_APP_KEY"),
		"alpha-vantage.p.rapidapi.com",
		os.Getenv("STOCK_APP_SERVER_URL"),
	)
	tradeManager := service.NewStockTradeHistoryManagementService(
		&repository,
		4,
		time.Second,
		fetcher,
	)
	adminService := service.NewAdminService(&tradeManager, &stockService)

	routers.NewStockRoutes(r, stockService)
	routers.NewAdminRoutes(r, &adminService)

	r.Run() // listen and serve on 0.0.0.0:8080
}
