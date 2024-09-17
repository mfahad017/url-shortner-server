package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.Engine) {

	v1 := g.Group("/v1")
	{
		RegisterAuthRoutes(v1)
		RegisterURLRoutes(v1)
	}

}
