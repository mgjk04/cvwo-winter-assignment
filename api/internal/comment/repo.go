package comment

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, t *Comment) (uuid.UUID, error)
	ReadByID(ctx context.Context, id uuid.UUID) (*Comment, error)
	ReadByPostID(ctx context.Context, postID uuid.UUID) ([]*Comment, error)
	UpdateByID(ctx context.Context, t *Comment) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

type commentRepo struct {
	db *pgxpool.Pool
}

func (r *commentRepo) Create(ctx context.Context, c *Comment) (uuid.UUID, error) {
	query := `INSERT INTO comments (content, post_id, author_id) VALUES ($1, $2, $3) RETURNING id`
	pool := r.db
	var id uuid.UUID
	err := pool.QueryRow(ctx, query, c.Content, c.PostID, c.AuthorID).Scan(&id)
	return id, err
}

func (r *commentRepo) ReadByID(ctx context.Context, id uuid.UUID) (*Comment, error) {
	query := `SELECT content, created_at, post_id, author_id FROM comments WHERE id=$1`
	pool := r.db
	p := &Comment{ID: id} 
	err := pool.QueryRow(ctx, query, id).Scan(&p.Content, &p.CreatedAt, &p.PostID, &p.AuthorID)
	return p, err
}

func (r *commentRepo) ReadByPostID(ctx context.Context, postID uuid.UUID) ([]*Comment, error) {
	query := `SELECT id, content, created_at, author_id FROM comments WHERE post_id=$1`
	pool := r.db
	comments := []*Comment{}
	rows, err := pool.Query(ctx, query, postID)
	for rows.Next() {
		c := &Comment{}
		err := rows.Scan(&c.ID, &c.Content, &c.CreatedAt, &c.AuthorID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, err
}

func (r *commentRepo) UpdateByID(ctx context.Context, c *Comment) error {
	query := `UPDATE comments SET content=$2, author_id=$3 WHERE id=$1`
	pool := r.db
	_, err := pool.Exec(ctx, query, c.ID, c.Content, c.AuthorID)
	return err
}

func (r *commentRepo) DeleteByID(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM comments WHERE id=$1`
	pool := r.db
	_, err := pool.Exec(ctx, query, id)
	return err
}

func NewCommentRepo(db *pgxpool.Pool) *commentRepo {
	return &commentRepo{db: db}
}