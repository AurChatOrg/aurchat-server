package service

import (
	"github.com/AurChatOrg/aurchat-server/internal/code"
	"github.com/AurChatOrg/aurchat-server/internal/model"
	"github.com/AurChatOrg/aurchat-server/internal/pkg/hasher"
	"github.com/AurChatOrg/aurchat-server/internal/pkg/token"
	"github.com/AurChatOrg/aurchat-server/internal/router/api/auth/repository"
	"github.com/bwmarrin/snowflake"
)

var (
	ErrAccountnameOrPassword = NewAuthError(
		code.ErrorAccountNameOrPassword,
		code.GetMessage(code.ErrorAccountNameOrPassword),
		nil,
	)

	ErrServerUnknown = NewAuthError(
		code.ServerUnknownError,
		code.GetMessage(code.ServerUnknownError),
		nil,
	)
)

type AuthService interface {
	SignIn(username, password string) (string, error)
	SignUp(username, password, email string) (string, error)
	ValidateToken(token string) (*model.User, error)
}

type authService struct {
	userRepo    repository.UserRepository
	hasher      *hasher.Hasher
	tokenGen    *token.Token
	idGenerator *snowflake.Node
}

func NewAuthService(
	userRepo repository.UserRepository,
	hasher *hasher.Hasher,
	tokenGen *token.Token,
	node *snowflake.Node,
) AuthService {
	return &authService{
		userRepo:    userRepo,
		hasher:      hasher,
		tokenGen:    tokenGen,
		idGenerator: node,
	}
}

func (s *authService) SignIn(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", ErrAccountnameOrPassword
	}

	if match, err := s.hasher.VerifyHash(password, user.Password); err != nil || !match {
		if !match && err == nil {
			return "", ErrAccountnameOrPassword
		} else {
			return "", ErrServerUnknown
		}
	}

	return s.tokenGen.Generate(user.Username, user.UserID)
}

func (s *authService) SignUp(username, password, email string) (string, error) {
	user := &model.User{
		UserID:   s.idGenerator.Generate().Int64(),
		Username: username,
		Email:    email,
		Password: s.hasher.Hash(password),
	}

	if err := s.userRepo.Create(user); err != nil {
		return "", err
	}

	return s.tokenGen.Generate(user.Username, user.UserID)
}

func (s *authService) ValidateToken(tokenStr string) (*model.User, error) {
	claims, err := s.tokenGen.Parse(tokenStr)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByUsername(claims.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
