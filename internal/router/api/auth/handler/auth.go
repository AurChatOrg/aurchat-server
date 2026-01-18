package handler

import (
	"net/http"

	"github.com/AurChatOrg/aurchat-server/internal/code"
	"github.com/AurChatOrg/aurchat-server/internal/dto"
	"github.com/AurChatOrg/aurchat-server/internal/router/api/auth/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// SignIn godoc
// @Summary      Sign in
// @Description  Sign in to the account using the username and password, and return a token if they are correct
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body body dto.SignInReq true "Sign in information"
// @Success      200  {object}  dto.TokenResp
// @Failure      400  {object}  dto.ErrorResp
// @Failure      500  {object}  dto.ErrorResp
// @Router       /auth/signIn [post]
func (h *AuthHandler) SignIn(c *gin.Context) {
	var req dto.SignInReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResp{
			Code: code.InvalidParameter,
		})
		return
	}

	token, err := h.authService.SignIn(req.Username, req.Password)
	if err != nil {
		if authErr, ok := err.(*service.AuthError); ok {
			c.JSON(http.StatusBadRequest, dto.ErrorResp{
				Code: authErr.Code,
			})
		} else {
			c.JSON(http.StatusInternalServerError, dto.ErrorResp{
				Code: code.ServerUnknownError,
			})
		}
		return
	}

	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		"token", // name
		token,   // value
		3600*24, // max-age
		"/",     // path
		"",      // domain
		true,    // secure
		true,    // httpOnly
	)

	c.JSON(http.StatusOK, dto.TokenResp{Token: token})
}

// SignUp godoc
// @Summary      Sign up
// @Description  Register an account using a username, email and password, and return a token if no errors occur
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body body dto.SignUpReq true "Sign up information"
// @Success      200  {object}  dto.TokenResp
// @Failure      400  {object}  dto.ErrorResp
// @Failure      500  {object}  dto.ErrorResp
// @Router       /auth/signUp [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	var req dto.SignUpReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResp{
			Code: code.InvalidParameter,
		})
		return
	}

	token, err := h.authService.SignUp(req.Username, req.Password, req.Email)
	if err != nil {
		if authErr, ok := err.(*service.AuthError); ok {
			c.JSON(http.StatusBadRequest, dto.ErrorResp{
				Code: authErr.Code,
			})
		} else {
			c.JSON(http.StatusInternalServerError, dto.ErrorResp{
				Code: code.ServerUnknownError,
			})
		}
		return
	}

	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(
		"token", // name
		token,   // value
		3600*24, // max-age
		"/",     // path
		"",      // domain
		true,    // secure
		true,    // httpOnly
	)

	c.JSON(http.StatusOK, dto.TokenResp{Token: token})
}
