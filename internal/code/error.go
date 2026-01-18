package code

// Global module
const (
	Success            = 0
	InvalidParameter   = 1001
	Unauthorized       = 1002
	Forbidden          = 1003
	NotFound           = 1004
	MethodNotAllowed   = 1005
	RequestTimeout     = 1006
	TooManyRequests    = 1007
	InternalError      = 1008
	ServiceUnavailable = 1009
)

// Auth module error code (2000-2099)
const (
	TokenExpired                        = 2001
	TokenInvalid                        = 2002
	AccountBanned                       = 2003
	ErrorAccountNameOrPassword          = 2004
	UserNotFound                        = 2005
	UserAlreadyExistsOrEmailAlreadyUsed = 2006
	EmailNotVerified                    = 2007
	PasswordTooWeak                     = 2008
	InvalidEmailFormat                  = 2009
	InvalidUsernameFormat               = 2010
)

// User module error code (2100-2199)
const (
	UserProfileUpdateFailed = 2101
	UserAvatarUploadFailed  = 2102
	UserNotActive           = 2103
	UserPermissionDenied    = 2104
)

// Chat module error code (2200-2299)
const (
	ChatRoomNotFound         = 2201
	ChatRoomPermissionDenied = 2202
	MessageSendFailed        = 2203
	MessageNotFound          = 2204
	ChatRoomFull             = 2205
)

// File module error code (2300-2399)
const (
	FileUploadFailed   = 2301
	FileSizeExceeded   = 2302
	FileTypeNotAllowed = 2303
	FileNotFound       = 2304
)

// Server module error code (9000-9999)
const (
	ServerUnknownError   = 9001
	DatabaseError        = 9002
	CacheError           = 9003
	ExternalServiceError = 9004
	ConfigurationError   = 9005
)

func GetMessage(code int) string {
	messages := map[int]string{
		Success:          "Success",
		InvalidParameter: "Invalid parameter",
		Unauthorized:     "Unauthorized",
		Forbidden:        "Forbidden",
		NotFound:         "Resource not found",

		TokenExpired:                        "Token expired",
		TokenInvalid:                        "Invalid token",
		AccountBanned:                       "Account is banned",
		ErrorAccountNameOrPassword:          "Incorrect username or password",
		UserNotFound:                        "User not found",
		UserAlreadyExistsOrEmailAlreadyUsed: "Username or email already exists",
		EmailNotVerified:                    "Email not verified",
		PasswordTooWeak:                     "Password is too weak",
		InvalidEmailFormat:                  "Invalid email format",
		InvalidUsernameFormat:               "Invalid username format",

		ServerUnknownError: "Internal server error",
		DatabaseError:      "Database error",
		CacheError:         "Cache error",
	}

	if msg, ok := messages[code]; ok {
		return msg
	}
	return "Unknown error"
}

func IsAuthError(code int) bool {
	return code >= 2000 && code < 2100
}

func IsServerError(code int) bool {
	return code >= 9000 && code < 10000
}

func IsClientError(code int) bool {
	return code >= 1000 && code < 2000
}
