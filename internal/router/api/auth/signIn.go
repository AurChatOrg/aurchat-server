package auth

import (
	"errors"
	"net/http"
	"runtime"

	"github.com/AurChatOrg/aurchat-server/internal/code"
	"github.com/AurChatOrg/aurchat-server/internal/config"
	"github.com/AurChatOrg/aurchat-server/internal/dto"
	"github.com/AurChatOrg/aurchat-server/internal/model"
	"github.com/AurChatOrg/aurchat-server/internal/pkg/hash"
	"github.com/AurChatOrg/aurchat-server/internal/pkg/token"
	"github.com/AurChatOrg/aurchat-server/internal/repo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SignIn godoc
// @Summary  Sign in
// @Description Sign in to the account using the username and password, and return a token if they are correct
// @Tags     Auth
// @Accept   json
// @Produce  json
// @Param	 body body		dto.SignInReq true "Sign in information"
// @Success  200  {object}  dto.TokenResp
// @Failure	 400  {object}  dto.ErrorResp
// @Router   /auth/sign-in [post]
func SignIn(c *gin.Context) {
	var req *dto.SignInReq

	if err := c.ShouldBindJSON(&req); /* Bind request content */ err != nil { // If an error occurs
		c.JSON(http.StatusBadRequest, dto.ErrorResp{
			Code: code.InvalidParameter, // Invalid parameter
		})
		return
	}

	// Find User
	var user model.User
	if err := repo.Postgres.First(&user, "name = ?", req.Username).Error; err != nil { // Find the user with the name <req.Username>
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, dto.ErrorResp{Code: code.UserNotFound})
		} else {
			c.JSON(http.StatusInternalServerError, dto.ErrorResp{Code: code.ServerUnknownError})
		}
		return
	}

	// Get User Password
	passwordHash := user.Password                                                  // Get User password hash from DB
	argon2Hash := hash.NewArgon2Hash(128*1024, 4, uint8(runtime.NumCPU()), 16, 32) // Create a new Argon2Hash instance

	match, err := argon2Hash.VerifyHash(req.Password, passwordHash) // Verify password
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResp{
			Code: code.ServerUnknownError,
		})
		return
	}

	if !match { // If password not right
		c.JSON(http.StatusBadRequest, dto.ErrorResp{Code: code.ErrorAccountNameOrPassword})
		return
	}

	// If password is right, generate a token and return
	brancaToken := token.NewBrancaToken(config.Cfg, 3600*24*30) // Create a new BrancaToken instance
	newToken := brancaToken.GenerateToken(req.Username, user.UserID)
	if newToken == "" {
		c.JSON(http.StatusInternalServerError, dto.ErrorResp{Code: code.ServerUnknownError})
		return
	}

	c.JSON(http.StatusOK, dto.TokenResp{Token: newToken})
}
