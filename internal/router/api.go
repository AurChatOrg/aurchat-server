// Package router defines RESTful route groups and their handlers.
package router

import (
	"github.com/AurChatOrg/aurchat-server/internal/config"
	"github.com/AurChatOrg/aurchat-server/internal/router/api/auth"
	"github.com/gin-gonic/gin"
)

// RegisterAPI mounts all /api/v1 endpoints.
func RegisterAPI(route *gin.Engine, cfg *config.Config) {
	api := route.Group("/api/v1")
	{
		// Health
		api.GET("/ping", ping)

		// Auth API
		auth.RegisterAuthAPI(api)
	}
}

// ping godoc
// @Summary  Health check
// @Description Returns pong to verify service liveness
// @Tags     Health check
// @Accept   json
// @Produce  json
// @Success  200  {object}	dto.Pong
// @Router   /ping [get]
func ping(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "pong"})
}
