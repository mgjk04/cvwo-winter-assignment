package post

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, t *Post) (uuid.UUID, error)
	ReadByID(ctx context.Context, id uuid.UUID) (*Post, error)
	ReadByTopicID(ctx context.Context, topicID uuid.UUID) ([]*Post, error)
	UpdateByID(ctx context.Context, t *Post) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

type postRepo struct {
	db *pgxpool.Pool
}

func (r *postRepo) Create(ctx context.Context, p *Post) (uuid.UUID, error) {
	query := `INSERT INTO posts (title, description, topic_id, author_id) VALUES ($1, $2, $3, $4) RETURNING id`
	pool := r.db
	var id uuid.UUID
	err := pool.QueryRow(ctx, query, p.Title, p.Description, p.TopicID, p.AuthorID).Scan(&id)
	return id, err
}

func (r *postRepo) ReadByID(ctx context.Context, id uuid.UUID) (*Post, error) {
	query := `SELECT title, description, topic_id, author_id FROM posts WHERE id=$1`
	pool := r.db
	p := &Post{ID: id} 
	err := pool.QueryRow(ctx, query, id).Scan(&p.Title, &p.Description, &p.TopicID, &p.AuthorID)
	return p, err
}

func (r *postRepo) ReadByTopicID(ctx context.Context, topicID uuid.UUID) ([]*Post, error) {
	query := `SELECT id, title, description, created_at, topic_id, author_id FROM posts WHERE topic_id=$1`
	pool := r.db
	posts := []*Post{}
	rows, err := pool.Query(ctx, query, topicID)
	for rows.Next() {
		p := &Post{}
		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.CreatedAt, &p.TopicID, &p.AuthorID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, err
}

func (r *postRepo) UpdateByID(ctx context.Context, p *Post) error {
	query := `UPDATE posts SET title=$2, description=$3, topic_id=$4, author_id=$5 WHERE id=$1`
	pool := r.db
	_, err := pool.Exec(ctx, query, p.ID, p.Title, p.Description, p.TopicID, p.AuthorID)
	return err
}

func (r *postRepo) DeleteByID(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM posts WHERE id=$1`
	pool := r.db
	_, err := pool.Exec(ctx, query, id)
	return err
}

func NewPostRepo(db *pgxpool.Pool) *postRepo {
	return &postRepo{db: db}
}