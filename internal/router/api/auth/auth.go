package auth

import "github.com/gin-gonic/gin"

// RegisterAuthAPI Register Auth API
func RegisterAuthAPI(route *gin.RouterGroup) {
	auth := route.Group("/auth")
	{
		auth.POST("/sign-in", SignIn)
		auth.POST("/sign-up", SignUp)
	}
}
