package routers

import (
	"github.com/ashishkujoy/paper-trading-backend/internal"
	"github.com/gin-gonic/gin"
)

func NewStockRoutes(router *gin.Engine, stocksService internal.StockService) *gin.RouterGroup {
	stocksRouter := router.Group("/stocks")
	stocksRouter.GET("/:symbol", func(ginCtx *gin.Context) {
		stocksService.GetBySymbol(ginCtx.Param("symbol"))
	})
	stocksRouter.GET("/gist", func(ginCtx *gin.Context) {
		gist, err := stocksService.GetStocksGist()
		if err != nil {
			ginCtx.JSON(500, gin.H{"message": "Error while fetching stocks gist"})
			return
		}
		ginCtx.JSON(200, gist)
	})
	return stocksRouter
}
