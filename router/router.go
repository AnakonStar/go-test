package router

import (
	"github.com/AnakonStar/go-api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/users", handlers.GetUsersMySQL)
	r.GET("/users/:id", handlers.GetUserMySQL)
	r.POST("/users", handlers.CreateUserMySQL)
	r.PUT("/users/:id", handlers.UpdateUserMySQL)
	r.DELETE("/users/:id", handlers.DeleteUserMySQL)

	return r
}
