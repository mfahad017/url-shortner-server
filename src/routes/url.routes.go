package routes

import (
	"server/src/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterURLRoutes(r *gin.RouterGroup) {

	var URLController = new(controllers.URLController)

	urlRoutes := r.Group("/url")
	{
		urlRoutes.POST("/", URLController.CreateURL)
		urlRoutes.GET("/user/:userId", URLController.ListURLsByUser)
		urlRoutes.GET("/:id", URLController.GetURL)
		urlRoutes.DELETE("/:id", URLController.DeleteURL)
		urlRoutes.PUT("/:id", URLController.UpdateURL)

	}

}
