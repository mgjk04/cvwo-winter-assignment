package post

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	CreatePost(ctx context.Context, p *Post) (uuid.UUID, error)
	UpdatePost(ctx context.Context, p *Post) error
	FindPosts(ctx context.Context, topicID uuid.UUID, page int, pageSize int) ([]*Post, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Post, error)
	DeletePost(ctx context.Context, id uuid.UUID) error
}

type postSvc struct {
	r Repo 
	//future auth service can be added later
}

func (s *postSvc) CreatePost(ctx context.Context, p *Post) (uuid.UUID, error) {
	// this thing is in charge of assembbling the post object to send to repo
	return s.r.Create(ctx, p);
}

func (s *postSvc) UpdatePost(ctx context.Context, t *Post) error {
	return s.r.UpdateByID(ctx, t)
}

func (s *postSvc) FindPosts(ctx context.Context, topicID uuid.UUID, page int, limit int) ([]*Post, error) {
	return s.r.ReadByTopicID(ctx, topicID, page, limit)
}

func (s *postSvc) FindByID(ctx context.Context, id uuid.UUID) (*Post, error) {
	return s.r.ReadByID(ctx, id)
}

func (s *postSvc) DeletePost(ctx context.Context, id uuid.UUID) error {
	return s.r.DeleteByID(ctx, id)
}

func NewPostSvc(r Repo) *postSvc {
	return &postSvc{r: r}
}