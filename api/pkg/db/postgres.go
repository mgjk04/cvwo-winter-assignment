package db

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

//error declarations
//TODO: move to dedicated file if big, including error handling logic
var (
	ErrDBConn = errors.New("database connection error")
	ErrDBconfig = errors.New("database configuration error")
)

func NewPool(ctx context.Context, url string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(url)
    if err != nil {
        return nil, errors.Join(ErrDBconfig, err)
    }
	dbpool, err := pgxpool.NewWithConfig(ctx, config)
	err = dbpool.Ping(ctx)
	if err != nil {
		return nil, errors.Join(ErrDBConn, err)
	}
	return dbpool, err
}