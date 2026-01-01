package comment

import (
	"errors"
	"github.com/jackc/pgx/v5"
)

var (
	ErrCommentNotFound = errors.New("comment not found")
	ErrUncaught	 = errors.New("uncaught error")
)
//TODO: refactor this below
func HandleError(err error) error {
	if err == nil {
		return nil
	}
	// var pgErr *pgconn.PgError
	// if errors.As(err, &pgErr) {
	// 	switch pgErr.Code {
	// 		//TODO: add the rest of the cases later
	// 	}
	// }
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrCommentNotFound
	}
	//default case, just return the original error wrapped with generic
	return errors.Join(ErrUncaught, err)
}
