package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/generalErrors"
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
	query := `INSERT INTO users (username) VALUES ($1) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(ctx, query, u.Username).Scan(&id)
	return id, generalErrors.PostgresqlErrorMap(err)
}

func (r *userRepo) ReadByUsername(ctx context.Context, username string) (*User, error) {
	query := `SELECT id, created_at, deleted_at FROM users WHERE username=$1 AND deleted_at IS NULL`
	u := &User{Username: username}
	err := r.db.QueryRow(ctx, query, username).Scan(&u.ID, &u.CreatedAt, &u.DeletedAt)
	return u, generalErrors.PostgresqlErrorMap(err)
}

func (r *userRepo) ReadByID(ctx context.Context, id uuid.UUID) (*User, error) {
	query := `SELECT username, created_at, deleted_at FROM users WHERE id=$1 AND deleted_at IS NULL`
	u := &User{ID: id}
	err := r.db.QueryRow(ctx, query, id).Scan(&u.Username, &u.CreatedAt, &u.DeletedAt)
	if err != nil {
		return nil, generalErrors.PostgresqlErrorMap(err)
	}
	return u, err
}

func (r *userRepo) DeleteByID(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE users SET deleted_at=NOW() WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return generalErrors.PostgresqlErrorMap(err)
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{db: db}
}