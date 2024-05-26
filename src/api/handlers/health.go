package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hailsayan/golang-api/api/helper"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Working!", true, 0))
	return
}
