package main

import (
	"server/src/config"
	c "server/src/config"
	"server/src/database"
	"server/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	c.LoadConfig()

	// TODO Connect with DB

	err := database.ConnectToDB()

	if err != nil {
		panic("Failed to connect to database")
	}

	app := gin.Default()

	if config.GetConfig().ENV == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	routes.RegisterRoutes(app)

	// c.ConfigureCors(app)

	app.Run(":" + c.GetConfig().Port)

	app.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Route '" + c.Request.URL.Path + "' not found",
		})
	})

}
