package errors

const (
	ErrPostNotFound = "post not found"

	ErrParseFile  = "error parse file"
	ErrDecodeJSON = "error decode json"
	ErrDecodeToml = "error decode toml"
	ErrJWTSign    = "error jwt sign"
	ErrHash       = "error hash"
	ErrWrongType  = "error wrong type"

	ErrSessionNotAuthenticated = "session not authenticated"

	ErrUserAlreadyExists = "user already exists"
	ErrUserNotFound      = "user not found"
	ErrUserLogin         = "wrong password or user doesn't exist"
	ErrBadToken          = "token is invalid"
	ErrExpiredToken      = "token is expired"

	ErrCommentNotFound = "comment not found"

	Success = "success"
)
