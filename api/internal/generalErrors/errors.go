package generalErrors

import (
	"errors"
)
//we define general errors here that error handling middleware and handlers, services and repos will use
var (
	ErrNotFound = errors.New("resource not found")
	ErrConflict = errors.New("resource conflict") //error caused by request data violating a constraint but is valid
	ErrInvalid = errors.New("request invalid") //error caused by request data that is not in an accepted format
	ErrUnauthorized = errors.New("unauthenticated")
	ErrForbidden = errors.New("unauthorized")
	ErrInternal = errors.New("internal error") //any other unhandled error
)
