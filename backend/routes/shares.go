// routes/shares.go
package routes

import (
	"peerpay/backend/controllers"

	"github.com/gin-gonic/gin"
)

func SharesRoutes(router *gin.Engine) {
	shares := router.Group("/shares")
	{
		shares.POST("/buy", controllers.BuyShares)
		shares.POST("/sell", controllers.SellShares)
		shares.GET("/:userID", controllers.GetShares)
	}
}
