package routers

import (
	"fmt"
	"github.com/ashishkujoy/paper-trading-backend/internal"
	"github.com/gin-gonic/gin"
)

func NewAdminRoutes(router *gin.Engine, adminService *internal.AdminService) *gin.RouterGroup {
	adminApi := router.Group("/admin")
	adminApi.POST("/update-stocks-trade-detail", func(ginCtx *gin.Context) {
		go func() {
			err := adminService.UpdateStockTradeDetails()
			if err != nil {
				fmt.Printf("Error while updating trade history: %v\n", err)
			}
		}()
		ginCtx.JSON(200, gin.H{"message": "Success"})
	})
	adminApi.POST("/stocks/symbol", func(ginCtx *gin.Context) {
		var request AddNewSymbolRequest
		err := ginCtx.BindJSON(&request)
		if err != nil {
			ginCtx.JSON(400, gin.H{"message": "Missing symbol"})
			return
		}
		err = adminService.AddNewSymbol(request.Symbol)
		if err != nil {
			fmt.Printf("Error while saving new symbol %v\n", err)
			ginCtx.JSON(500, gin.H{"message": "Failed to save new symbol"})
		}
	})
	return adminApi
}
