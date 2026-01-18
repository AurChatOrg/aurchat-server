package token

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/AurChatOrg/aurchat-server/internal/code"
	"github.com/AurChatOrg/aurchat-server/internal/config"
	"github.com/AurChatOrg/aurchat-server/internal/pkg/logger"
	"github.com/essentialkaos/branca/v2"
	"go.uber.org/zap"
)

// Token Branca's encapsulation structure
type Token struct {
	branca branca.Branca
	ttl    uint32
}

// UserClaims User information
type UserClaims struct {
	Username string `json:"username"`
	UserID   int64  `json:"userID"`
}

// NewToken New branca instance
func NewToken(cfg *config.Config, ttl uint32) *Token {
	brc, err := branca.NewBranca([]byte(cfg.Auth.Keys)) // Create new branca struct

	if err != nil {
		logger.Logger.Fatal("Error creating Branca struct", zap.Error(err))
	}

	return &Token{branca: brc, ttl: ttl}
}

// Generate Generate a new branca token
func (b *Token) Generate(username string, userId int64) (string, error) {
	if username == "" {
		return "", errors.New("username cannot be blank")
	}

	// Encode User information
	userInfo := &UserClaims{
		Username: username,
		UserID:   userId,
	}

	userData, err := json.Marshal(userInfo)
	if err != nil {
		logger.Logger.Error("Error marshalling user info", zap.Error(err))
		return "", err
	}

	// Generate branca token
	token, err := b.branca.EncodeToString([]byte(userData))
	if err != nil {
		logger.Logger.Error("Error encoding token", zap.Error(err))
		return "", err
	}

	return token, nil
}

// Parse Parse branca token
func (b *Token) Parse(token string) (UserClaims, error) {
	// Decode token
	raw, err := b.branca.DecodeString(token)
	if err != nil {
		logger.Logger.Error("Error decoding token", zap.Error(err))
		return UserClaims{}, err
	}

	// Decode JSON
	var userInfo UserClaims
	if err = json.Unmarshal([]byte(raw.Payload()), &userInfo); err != nil {
		logger.Logger.Error("Error unmarshalling token", zap.Error(err))
		return UserClaims{}, err
	}

	// Expired inspection
	expired := raw.IsExpired(b.ttl)
	if expired {
		logger.Logger.Error("Token is expired", zap.String("token", token))
		return UserClaims{}, errors.New(strconv.Itoa(code.TokenExpired))
	}

	// Return User information
	return userInfo, nil
}
