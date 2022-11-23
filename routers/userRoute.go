package routers

import (
	"github.com/ashishkujoy/paper-trading-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func NewUserRoutes(router *gin.Engine, userService *service.UserService) *gin.RouterGroup {
	usersRoute := router.Group("/users")

	usersRoute.GET("/:userId/portfolio", func(ginCtx *gin.Context) {
		userId := ginCtx.Param("userId")
		portfolioView, err := userService.GetPortfolioForUser(userId)
		if err != nil {
			ginCtx.JSON(500, gin.H{"message": "error while fetching portfolio"})
			return
		}
		ginCtx.JSON(200, portfolioView)
	})

	usersRoute.POST("/", func(ginCtx *gin.Context) {
		request := CreateNewUserRequest{}
		err := ginCtx.BindJSON(&request)
		if err != nil {
			ginCtx.JSON(400, gin.H{"message": "Malformed create new user request"})
			return
		}

		portfolio, err := userService.CreateNewUser(request.Username, request.Balance)
		if err != nil {
			ginCtx.JSON(500, gin.H{"message": "Error while creating new user"})
			return
		}

		ginCtx.JSON(200, portfolio)
	})

	//usersRoute.POST("/:userId/portfolio/stock", func(ginCtx *gin.Context) {
	//	request := StockBuySellRequest{}
	//	err := ginCtx.BindJSON(&request)
	//	if err != nil {
	//		ginCtx.JSON(400, gin.H{"message": "Malformed buy/sell stocks request"})
	//		return
	//	}
	//
	//	userId := ginCtx.Param("userId")
	//
	//	userService.ExecuteOrder(request.Action, request.Quantity, request.Symbol)
	//})

	return usersRoute
}
