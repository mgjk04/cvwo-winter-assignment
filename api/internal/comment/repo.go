package comment

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/generalErrors"
)

type Repo interface {
	Create(ctx context.Context, t *Comment) (uuid.UUID, error)
	ReadByID(ctx context.Context, id uuid.UUID) (*Comment, error)
	ReadByPostID(ctx context.Context, postID uuid.UUID, page int, limit int) ([]*Comment, error)
	UpdateByID(ctx context.Context, t *Comment) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

type commentRepo struct {
	db *pgxpool.Pool
}

func (r *commentRepo) Create(ctx context.Context, c *Comment) (uuid.UUID, error) {
	query := `INSERT INTO comments (content, post_id, author_id) VALUES ($1, $2, $3) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(ctx, query, c.Content, c.PostID, c.AuthorID).Scan(&id)
	return id, generalErrors.PostgresqlErrorMap(err)
}

func (r *commentRepo) ReadByID(ctx context.Context, id uuid.UUID) (*Comment, error) {
	query := `SELECT content, comments.created_at, post_id, author_id, users.username FROM comments JOIN users ON author_id = users.id WHERE comments.id=$1`
	p := &Comment{ID: id} 
	err := r.db.QueryRow(ctx, query, id).Scan(&p.Content, &p.CreatedAt, &p.PostID, &p.AuthorID, &p.AuthorName)
	return p, generalErrors.PostgresqlErrorMap(err)
}

func (r *commentRepo) ReadByPostID(ctx context.Context, postID uuid.UUID, page int, limit int) ([]*Comment, error) {
	query := `SELECT comments.id, content, comments.created_at, post_id, author_id, users.username FROM comments JOIN users ON author_id = users.id WHERE post_id=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	rows, err := r.db.Query(ctx, query, postID, limit, (page - 1) * limit)
	if err != nil {
		return nil, generalErrors.PostgresqlErrorMap(err)
	}
	defer rows.Close()
	comments := []*Comment{}
	for rows.Next() {
		c := &Comment{}
		err := rows.Scan(&c.ID, &c.Content, &c.CreatedAt, &c.PostID, &c.AuthorID, &c.AuthorName)
		if err != nil {
			return nil, generalErrors.PostgresqlErrorMap(err)
		}
		comments = append(comments, c)
	}
	return comments, generalErrors.PostgresqlErrorMap(err)
}

func (r *commentRepo) UpdateByID(ctx context.Context, c *Comment) error {
	query := `UPDATE comments SET content=$2, author_id=$3 WHERE id=$1`
	pool := r.db
	_, err := pool.Exec(ctx, query, c.ID, c.Content, c.AuthorID)
	return generalErrors.PostgresqlErrorMap(err)
}

func (r *commentRepo) DeleteByID(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM comments WHERE id=$1`
	pool := r.db
	_, err := pool.Exec(ctx, query, id)
	return generalErrors.PostgresqlErrorMap(err)
}

func NewCommentRepo(db *pgxpool.Pool) *commentRepo {
	return &commentRepo{db: db}
}