// Package router defines RESTful route groups and their handlers.
package router

import (
	"github.com/AurChatOrg/aurchat-server/internal/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterAPI mounts all /api/v1 endpoints.
func RegisterAPI(route *gin.Engine, cfg *config.Config) {
	api := route.Group("/api/v1")
	{
		// Health
		api.GET("/ping", ping)

		// Swagger UI
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}

// ping godoc
// @Summary  Health check
// @Description Returns pong to verify service liveness
// @Tags     health
// @Accept   json
// @Produce  json
// @Success  200  {object}  map[string]string "{"msg":"pong"}"
// @Router   /ping [get]
func ping(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "pong"})
}
