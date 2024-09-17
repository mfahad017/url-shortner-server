package config

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigureCors(g *gin.Engine) {
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com", "https://bar.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://allowed.com"
		},
		MaxAge: 12 * time.Hour,
	}))
}
