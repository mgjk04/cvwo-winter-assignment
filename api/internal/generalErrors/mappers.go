package generalErrors

import (
	"errors"
	"log/slog"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

//implementation mappers
//takes an implementation specific error and converts it to general errors
//TODO: add logging of errors
func PostgresqlErrorMap(err error) error {
	if err == nil {
		return nil
	}
	slog.Error(err.Error());
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
			case "23505": //duplicate constraint violated
				return ErrConflict
			//TODO: add the rest of the cases later
		}
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}
	//default case
	return ErrInternal
}

//response 

