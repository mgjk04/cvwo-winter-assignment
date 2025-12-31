package user

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, u *User) (uuid.UUID, error)
	ReadByUsername(ctx context.Context, username string) (*User, error)
	ReadByID(ctx context.Context, id uuid.UUID) (*User, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

type userRepo struct {
	db *pgxpool.Pool
}

func (r *userRepo) Create(ctx context.Context, u *User) (uuid.UUID, error) {
	query := `INSERT INTO users (id, username) VALUES ($1, $2)`
	pool := r.db
	var id uuid.UUID
	err := pool.QueryRow(ctx, query, u.ID, u.Username).Scan(&id)
	return id, err
}

func (r *userRepo) ReadByUsername(ctx context.Context, username string) (*User, error) {
	query := `SELECT id, created_at, deleted_at FROM users WHERE username=$1 AND deleted_at IS NULL`
	pool := r.db
	u := &User{Username: username}
	err := pool.QueryRow(ctx, query, username).Scan(&u.ID, &u.CreatedAt, &u.DeletedAt)
	return u, err
}

func (r *userRepo) ReadByID(ctx context.Context, id uuid.UUID) (*User, error) {
	query := `SELECT username, created_at, deleted_at FROM users WHERE id=$1 AND deleted_at IS NULL`
	pool := r.db
	u := &User{ID: id}
	err := pool.QueryRow(ctx, query, id).Scan(&u.Username, &u.CreatedAt, &u.DeletedAt)
	return u, err
}

func (r *userRepo) DeleteByID(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE users SET deleted_at=NOW() WHERE id=$1`
	pool := r.db
	_, err := pool.Exec(ctx, query, id)
	return err
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{db: db}
}