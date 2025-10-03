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
	"gorm.io/gorm"
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
	errs := repo.Postgres.Transaction(func(tx *gorm.DB) error {
		/* SnowFlake instance */
		node, err := snowflake.NewNode(config.Cfg.Snowflake.WorkerID) // Create a SnowFlake Node
		if err != nil {
			logger.Logger.Error("Create a new Snowflake Node error", zap.Error(err))
			c.JSON(http.StatusInternalServerError, dto.ErrorResp{Code: code.ServerUnknownError})
			return err
		}

		/* Argon2Hash instance */
		argon2Hash := hash.NewArgon2Hash(config.Cfg.Hash.Memory, config.Cfg.Hash.Interactions, uint8(runtime.NumCPU()), config.Cfg.Hash.SaltLength, config.Cfg.Hash.KeyLength)

		/* Generate User ID and Password(hashed) */
		id := int64(node.Generate())                          // Generate User ID
		passwordHash := argon2Hash.GenerateHash(req.Password) // Hash password

		user := model.User{
			UserID:   id,
			Name:     req.Username,
			Password: passwordHash,
			Email:    req.Email,
		}

		tx.Create(&user)

		brancaToken := token.NewBrancaToken(config.Cfg, config.Cfg.Auth.TTL) // Create a new BrancaToken instance
		newToken, err := brancaToken.GenerateToken(user.Name, user.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResp{Code: code.ServerUnknownError})
			return err
		}

		c.JSON(http.StatusOK, dto.TokenResp{Token: newToken})
		return nil
	})

	if errs != nil {
		return
	}
}
