// routes/auth.go
package routes

import (
	"peerpay/backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/login", controllers.Login)
		auth.POST("/referral", controllers.Referral)
	}
}
