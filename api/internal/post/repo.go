package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/generalErrors"
)

type Repo interface {
	Create(ctx context.Context, t *Post) (uuid.UUID, error)
	ReadByID(ctx context.Context, id uuid.UUID) (*Post, error)
	ReadByTopicID(ctx context.Context, topicID uuid.UUID, page int, limit int) ([]*Post, error)
	UpdateByID(ctx context.Context, t *Post) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

type postRepo struct {
	db *pgxpool.Pool
}

func (r *postRepo) Create(ctx context.Context, p *Post) (uuid.UUID, error) {
	query := `INSERT INTO posts (title, description, topic_id, author_id) VALUES ($1, $2, $3, $4) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(ctx, query, p.Title, p.Description, p.TopicID, p.AuthorID).Scan(&id)
	return id, generalErrors.PostgresqlErrorMap(err)
}

func (r *postRepo) ReadByID(ctx context.Context, id uuid.UUID) (*Post, error) {
	query := `SELECT title, description, topic_id, author_id FROM posts WHERE id=$1`
	p := &Post{ID: id} 
	err := r.db.QueryRow(ctx, query, id).Scan(&p.Title, &p.Description, &p.TopicID, &p.AuthorID)
	return p, generalErrors.PostgresqlErrorMap(err)
}

func (r *postRepo) ReadByTopicID(ctx context.Context, topicID uuid.UUID, page int, limit int) ([]*Post, error) {
	query := `SELECT id, title, description, created_at, topic_id, author_id FROM posts WHERE topic_id=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	pool := r.db
	rows, err := pool.Query(ctx, query, topicID, limit, (page-1)*limit)
	if err != nil {
		return nil, generalErrors.PostgresqlErrorMap(err)
	}
	defer rows.Close()
	posts := []*Post{}
	for rows.Next() {
		p := &Post{}
		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.CreatedAt, &p.TopicID, &p.AuthorID)
		if err != nil {
			return nil, generalErrors.PostgresqlErrorMap(err)
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func (r *postRepo) UpdateByID(ctx context.Context, p *Post) error {
	query := `UPDATE posts SET title=$2, description=$3, topic_id=$4, author_id=$5 WHERE id=$1`
	pool := r.db
	_, err := pool.Exec(ctx, query, p.ID, p.Title, p.Description, p.TopicID, p.AuthorID)
	return generalErrors.PostgresqlErrorMap(err)
}

func (r *postRepo) DeleteByID(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM posts WHERE id=$1`
	pool := r.db
	_, err := pool.Exec(ctx, query, id)
	return generalErrors.PostgresqlErrorMap(err)
}

func NewPostRepo(db *pgxpool.Pool) *postRepo {
	return &postRepo{db: db}
}