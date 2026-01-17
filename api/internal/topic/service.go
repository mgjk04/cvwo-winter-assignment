package topic

import (
	"context"

	"github.com/google/uuid"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/generalErrors"
)

type Service interface {
	CreateTopic(ctx context.Context, t *Topic) (uuid.UUID, error)
	UpdateTopic(ctx context.Context, userID uuid.UUID, t *Topic) error
	FindTopics(ctx context.Context, page int, pageSize int) ([]*Topic, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Topic, error)
	DeleteTopic(ctx context.Context, userID uuid.UUID, topicID uuid.UUID) error
}

type topicSvc struct {
	r Repo 
}

func (s *topicSvc) CreateTopic(ctx context.Context, t *Topic) (uuid.UUID, error) {
	return s.r.Create(ctx, t);
}

func (s *topicSvc) UpdateTopic(ctx context.Context, userID uuid.UUID, t *Topic) error {
	isAuthor, err := s.verifyAuthor(ctx, userID, t.ID)
	if err != nil {
		return err
	}
	if !isAuthor {
		return generalErrors.ErrForbidden
	}
	return s.r.UpdateByID(ctx, t)
}

func (s *topicSvc) FindTopics(ctx context.Context, page int, pageSize int) ([]*Topic, error) {
	return s.r.ReadMany(ctx, page, pageSize)
}

func (s *topicSvc) FindByID(ctx context.Context, id uuid.UUID) (*Topic, error) {
	return s.r.ReadByID(ctx, id)
}

func (s *topicSvc) DeleteTopic(ctx context.Context, userID uuid.UUID, topicID uuid.UUID) error {
	isAuthor, err := s.verifyAuthor(ctx, userID, topicID)
	if err != nil {
		return err
	}
	if !isAuthor {
		return generalErrors.ErrForbidden
	}

	return s.r.DeleteByID(ctx, topicID)
}

func (s *topicSvc) verifyAuthor(ctx context.Context, userID uuid.UUID, topicID uuid.UUID) (bool, error) {
	topic, err := s.r.ReadByID(ctx, topicID)
	if err != nil {
		return false, err
	}
	return topic.AuthorID == userID, err
}

func NewTopicSvc(r Repo) *topicSvc {
	return &topicSvc{r: r}
}