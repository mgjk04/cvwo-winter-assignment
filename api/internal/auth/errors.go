package auth

import (
	"errors"
)
//these are more specific errors that are used by the auth middleware

var (
	ErrSecretNotFound = errors.New("secret inaccessible, check environment variables")
	ErrAccessTokenGen = errors.New("unable to generate access token")
	ErrRefreshTokenGen = errors.New("unable to generate refresh token")
	ErrTokenMalformed = errors.New("token malformed")
	ErrTokenSignatureInvalid = errors.New("token signature is invalid")
	ErrTokenExpired = errors.New("token expired, sign in again")
	ErrTokenNotValidYet = errors.New("token not yet valid, sign in or wait")
)