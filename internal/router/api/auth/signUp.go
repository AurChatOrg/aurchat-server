package auth

import (
	"net/http"
	"runtime"

	"github.com/AurChatOrg/aurchat-server/internal/code"
	"github.com/AurChatOrg/aurchat-server/internal/config"
	"github.com/AurChatOrg/aurchat-server/internal/dto"
	"github.com/AurChatOrg/aurchat-server/internal/model"
	"github.com/AurChatOrg/aurchat-server/internal/pkg/hash"
	"github.com/AurChatOrg/aurchat-server/internal/pkg/logger"
	"github.com/AurChatOrg/aurchat-server/internal/pkg/token"
	"github.com/AurChatOrg/aurchat-server/internal/repo"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SignUp godoc
// @Summary  Sign up
// @Description Register an account using a username and password, and return a token if no errors occur.
// @Tags     Auth
// @Accept   json
// @Produce  json
// @Param	 body body		dto.SignUpReq true "Sign up information"
// @Success  200  {object}  dto.TokenResp
// @Failure	 400  {object}  dto.ErrorResp
// @Router   /auth/sign-up [post]
func SignUp(c *gin.Context) {
	var req *dto.SignUpReq

	if err := c.ShouldBindJSON(&req); /* Bind request content */ err != nil { // If an error occurs
		c.JSON(http.StatusBadRequest, dto.ErrorResp{
			Code: code.InvalidParameter, // Invalid parameter
		})
		return
	}

	// Check if the username or email is already in use
	var count int64
	repo.Postgres.Model(&model.User{}).Where("name = ? OR email = ?", req.Username, req.Email).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, dto.ErrorResp{Code: code.UserAlreadyExistsOrEmailAlreadyUsed})
		return
	}

	// Create a new User struct
	/* SnowFlake instance */
	nodeId := int64(1)                     // Set Node ID
	node, err := snowflake.NewNode(nodeId) // Create a SnowFlake Node
	if err != nil {
		logger.Logger.Error("Create a new Snowflake Node error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, dto.ErrorResp{Code: code.ServerUnknownError})
		return
	}

	/* Argon2Hash instance */
	argon2Hash := hash.NewArgon2Hash(32*1024, 4, uint8(runtime.NumCPU()), 16, 32)

	/* Generate User ID and Password(hashed) */
	id := int64(node.Generate())                          // Generate User ID
	passwordHash := argon2Hash.GenerateHash(req.Password) // Hash password

	user := model.User{
		UserID:   id,
		Name:     req.Username,
		Password: passwordHash,
		Email:    req.Email,
	}

	repo.Postgres.Create(&user)

	brancaToken := token.NewBrancaToken(config.Cfg, 3600*24*30) // Create a new BrancaToken instance
	newToken := brancaToken.GenerateToken(user.Name, user.UserID)
	if newToken == "" {
		c.JSON(http.StatusInternalServerError, dto.ErrorResp{Code: code.ServerUnknownError})
		return
	}

	c.JSON(http.StatusOK, dto.TokenResp{Token: newToken})
}
