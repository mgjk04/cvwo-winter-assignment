package topic

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/google/uuid"
)

//NOTE: add tx.rollback if extension logic requires mutliple queries 
type Repo interface {
	Create(ctx context.Context, t *Topic) (uuid.UUID, error)
	ReadByID(ctx context.Context, id uuid.UUID) (*Topic, error)
	UpdateByID(ctx context.Context, t *Topic) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

type topicRepo struct {
	db *pgxpool.Pool
}

func (r *topicRepo) Create(ctx context.Context, t *Topic) (uuid.UUID, error) {
	query := `INSERT INTO topics (topicname, description, author_id) VALUES ($1, $2, $3) RETURNING id`
	pool := r.db
	var id uuid.UUID
	err := pool.QueryRow(ctx, query, t.TopicName, t.Description, t.AuthorID).Scan(&id)
	return id, err
}

func (r *topicRepo) ReadByID(ctx context.Context, id uuid.UUID) (*Topic, error) {
	query := `SELECT topicname, description, created_at, author_id FROM topics WHERE id=$1 AND deleted_at IS NULL`
	pool := r.db
	t := &Topic{ID: id}
	err := pool.QueryRow(ctx, query, id).Scan(&t.TopicName, &t.Description, &t.CreatedAt, &t.AuthorID)
	return t, err
}

func (r *topicRepo) UpdateByID(ctx context.Context, t *Topic) error {
	query := `UPDATE topics SET topicname=$2, description=$3, author_id=$5 WHERE id=$1`
	pool := r.db
	_, err := pool.Exec(ctx, query, t.ID, t.TopicName, t.Description, t.AuthorID)
	return err
}

func (r *topicRepo) DeleteByID(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM topics WHERE id=$1`
	pool := r.db
	_, err := pool.Exec(ctx, query, id)
	return err
}

func NewTopicRepo(db *pgxpool.Pool) *topicRepo {
	return &topicRepo{db: db}
}
