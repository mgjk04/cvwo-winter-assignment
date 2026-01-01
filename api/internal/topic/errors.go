package topic

import (
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrTopicNotFound = errors.New("topic not found")
	ErrTopicExists   = errors.New("topic already exists")
	ErrUncaught	 = errors.New("uncaught error")
)
//TODO: refactor this below
func HandleError(err error) error {
	if err == nil {
		return nil
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
			case "23505":
				return ErrTopicExists
			//TODO: add the rest of the cases later
		}
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrTopicNotFound
	}
	//default case, just return the original error wrapped with generic
	return errors.Join(ErrUncaught, err)
}
