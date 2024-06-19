// routes/admin.go
package routes

import (
	"peerpay/backend/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	admin := router.Group("/admin")
	{
		admin.GET("/users", controllers.GetUsers)
		admin.DELETE("/user/:userID", controllers.DeleteUser)
		admin.GET("/dormant-accounts", controllers.GetDormantAccounts)
	}
}
