package code

// Gloval module
const (
	InvalidParameter = 1001
)

// User module error code
const (
	TokenExpired                        = 2001
	AccountBanned                       = 2002
	ErrorAccountNameOrPassword          = 2003
	UserNotFound                        = 2004
	UserAlreadyExistsOrEmailAlreadyUsed = 2005
)

// Server module error code
const (
	ServerUnknownError = 9001
)
