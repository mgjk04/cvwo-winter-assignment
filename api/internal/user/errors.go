package user

import (
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user already exists")
	ErrUncaught	 = errors.New("uncaught error")
)
//takes an implementation specific error and converts it to our general errors above
//TODO: refactor this below
func HandleError(err error) error {
	if err == nil {
		return nil
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
			case "23505":
				return ErrUserExists
			//TODO: add the rest of the cases later
		}
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrUserNotFound
	}
	//default case, just return the original error wrapped with generic
	return errors.Join(ErrUncaught, err)
}
