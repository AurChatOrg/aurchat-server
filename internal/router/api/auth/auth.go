package auth

import (
	"runtime"

	"github.com/AurChatOrg/aurchat-server/internal/config"
	"github.com/AurChatOrg/aurchat-server/internal/pkg/hasher"
	"github.com/AurChatOrg/aurchat-server/internal/pkg/logger"
	"github.com/AurChatOrg/aurchat-server/internal/pkg/token"
	"github.com/AurChatOrg/aurchat-server/internal/repo"
	"github.com/AurChatOrg/aurchat-server/internal/router/api/auth/handler"
	"github.com/AurChatOrg/aurchat-server/internal/router/api/auth/repository"
	"github.com/AurChatOrg/aurchat-server/internal/router/api/auth/service"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RegisterAuthAPI Register Auth API
func RegisterAuthAPI(route *gin.RouterGroup) {
	hasher := hasher.NewHasher(config.Cfg.Hash.Memory, config.Cfg.Hash.Inerations, uint8(runtime.NumCPU()), config.Cfg.Hash.SaltLength, 32)
	tokenGen := token.NewToken(config.Cfg, 3600*24*30)
	node, err := snowflake.NewNode(config.Cfg.Snowflake.WorkerID) // Create a SnowFlake Node
	if err != nil {
		logger.Logger.Error("Create a new Snowflake Node error", zap.Error(err))
	}

	userRepo := repository.NewUserRepository(repo.Postgres)
	service := service.NewAuthService(userRepo, hasher, tokenGen, node)
	handler := handler.NewAuthHandler(service)

	auth := route.Group("/auth")
	{
		auth.POST("/signIn", handler.SignIn)
		auth.POST("/signUp", handler.SignUp)
	}
}
