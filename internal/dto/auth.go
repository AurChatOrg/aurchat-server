package dto

import "github.com/go-playground/validator/v10"

// SignInReq Sign in request structure
type SignInReq struct {
	Username string `json:"username" binding:"required,min=2,max=20" example:"xxx"`
	Password string `json:"password" binding:"required,min=2,max=32"       example:"******"`
	Email    string `json:"email" binding:"required"       example:"xxx@example.com"`
}

// SignUpReq Sign up request structure
type SignUpReq struct {
	Username string `json:"username" binding:"required,min=2,max=20" example:"xxx"`
	Password string `json:"password" binding:"required,min=2,max=32"       example:"******"`
	Email    string `json:"email" binding:"required"       example:"xxx@example.com"`
}

// TokenResp Sign in or Sign up response structure
type TokenResp struct {
	Token string `json:"token" example:"Token"`
}

// Validate Single case verifier
var Validate = validator.New()
