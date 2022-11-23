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
	stocksRepository := repository.NewStockRepositoryRedisImpl(redisClient)
	stockService := service.NewStockService(&stocksRepository)
	fetcher := service.NewStockPriceFetcher(
		os.Getenv("STOCK_APP_KEY"),
		"alpha-vantage.p.rapidapi.com",
		os.Getenv("STOCK_APP_SERVER_URL"),
	)
	tradeManager := service.NewStockTradeHistoryManagementService(
		&stocksRepository,
		4,
		time.Second,
		fetcher,
	)
	userRepository := repository.NewUserRepository(redisClient)
	adminService := service.NewAdminService(&tradeManager, &stockService)
	userService := service.NewUserService(&userRepository, &stocksRepository)

	routers.NewStockRoutes(r, stockService)
	routers.NewAdminRoutes(r, &adminService)
	routers.NewUserRoutes(r, &userService)

	r.Run() // listen and serve on 0.0.0.0:8080
}
