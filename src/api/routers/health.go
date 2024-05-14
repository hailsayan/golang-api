package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hailsayan/golang-api/api/handlers"
)

func Health (r *gin.RouterGroup){
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
	r.POST("/", handler.HealthPost)
	r.POST("/:id", handler.HealthPostById)
}