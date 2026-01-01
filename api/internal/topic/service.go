package topic

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	CreateTopic(ctx context.Context, t *Topic) (uuid.UUID, error)
	UpdateTopic(ctx context.Context, t *Topic) error
	FindTopics(ctx context.Context, page int, pageSize int) ([]*Topic, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Topic, error)
	DeleteTopic(ctx context.Context, id uuid.UUID) error
}

type topicSvc struct {
	r Repo 
	//future auth service can be added later
}

func (s *topicSvc) CreateTopic(ctx context.Context, t *Topic) (uuid.UUID, error) {
	// this thing is in charge of assembbling the user object to send to repo
	return s.r.Create(ctx, t);
}

func (s *topicSvc) UpdateTopic(ctx context.Context, t *Topic) error {
	return s.r.UpdateByID(ctx, t)
}

func (s *topicSvc) FindTopics(ctx context.Context, page int, pageSize int) ([]*Topic, error) {
	return s.r.ReadMany(ctx, page, pageSize)
}

func (s *topicSvc) FindByID(ctx context.Context, id uuid.UUID) (*Topic, error) {
	return s.r.ReadByID(ctx, id)
}

func (s *topicSvc) DeleteTopic(ctx context.Context, id uuid.UUID) error {
	return s.r.DeleteByID(ctx, id)
}

func NewTopicSvc(r Repo) *topicSvc {
	return &topicSvc{r: r}
}