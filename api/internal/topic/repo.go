package topic

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

//NOTE: add tx.rollback if extension logic requires mutliple queries
type Repo interface {
	Create(ctx context.Context, t *Topic) (uuid.UUID, error)
	ReadByID(ctx context.Context, id uuid.UUID) (*Topic, error)
	ReadMany(ctx context.Context, page int, limit int) ([]*Topic, error)
	UpdateByID(ctx context.Context, t *Topic) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

type topicRepo struct {
	db *pgxpool.Pool
}

//TODO: error handler
func (r *topicRepo) Create(ctx context.Context, t *Topic) (uuid.UUID, error) {
	query := `INSERT INTO topics (topicname, description, author_id) VALUES ($1, $2, $3) RETURNING id`
	var id uuid.UUID
	err := r.db.QueryRow(ctx, query, t.TopicName, t.Description, t.AuthorID).Scan(&id)
	return id, HandleError(err)
}

func (r *topicRepo) ReadMany(ctx context.Context, page int, limit int) ([]*Topic, error) {
	query := `SELECT id, topicname, description, created_at, author_id FROM topics ORDER BY created_at DESC LIMIT $1 OFFSET $2 `
	rows, err := r.db.Query(ctx, query, limit, (page - 1) * limit)
	if err != nil {
		return nil, HandleError(err)
	}
	defer rows.Close()
	topics := []*Topic{}
	for rows.Next() {
		t := &Topic{}
		err := rows.Scan(&t.ID, &t.TopicName, &t.Description, &t.CreatedAt, &t.AuthorID)
		if err != nil {
			return nil, HandleError(err)
		}
		topics = append(topics, t)
	}
	return topics, nil
}

func (r *topicRepo) ReadByID(ctx context.Context, id uuid.UUID) (*Topic, error) {
	query := `SELECT topicname, description, created_at, author_id FROM topics WHERE id=$1`
	t := &Topic{ID: id}
	err := r.db.QueryRow(ctx, query, id).Scan(&t.TopicName, &t.Description, &t.CreatedAt, &t.AuthorID)
	return t, HandleError(err)
}

func (r *topicRepo) UpdateByID(ctx context.Context, t *Topic) error {
	query := `UPDATE topics SET topicname=$2, description=$3, author_id=$4 WHERE id=$1`
	_, err := r.db.Exec(ctx, query, t.ID, t.TopicName, t.Description, t.AuthorID)
	return HandleError(err)
}

func (r *topicRepo) DeleteByID(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM topics WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return HandleError(err)
}

func NewTopicRepo(db *pgxpool.Pool) *topicRepo {
	return &topicRepo{db: db}
}
