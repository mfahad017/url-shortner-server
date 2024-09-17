package routes

import (
	"server/src/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.RouterGroup) {

	var AuthController = new(controllers.AuthController)

	authRoute := r.Group("/auth")
	{
		authRoute.POST("/signup", AuthController.SignUp)
		authRoute.POST("/login", AuthController.Login)
		// TODO Rest Password
		// TODO Refresh Token
		// TODO forgot Password
	}

}
