package handlers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *UserHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{

		api.GET("/users/:id", userHandler.GetUser)
		api.POST("/users", userHandler.CreateUser)
		api.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Route ทำงานปกติ"})
		})
	}

	return r
}
