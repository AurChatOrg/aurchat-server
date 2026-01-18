package service

import "github.com/AurChatOrg/aurchat-server/internal/code"

type AuthError struct {
	Code    int
	Message string
	Err     error
}

func (e *AuthError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return code.GetMessage(e.Code)
}

func NewAuthError(code int, message string, err error) *AuthError {
	return &AuthError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
