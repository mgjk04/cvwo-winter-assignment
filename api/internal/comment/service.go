package comment

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	CreateComment(ctx context.Context, p *Comment) (uuid.UUID, error)
	UpdateComment(ctx context.Context, p *Comment) error
	FindComments(ctx context.Context, topicID uuid.UUID, page int, limit int) ([]*Comment, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Comment, error)
	DeleteComment(ctx context.Context, id uuid.UUID) error
}

type commentSvc struct {
	r Repo 
	//future auth service can be added later
}

func (s *commentSvc) CreateComment(ctx context.Context, c *Comment) (uuid.UUID, error) {
	// this thing is in charge of assembbling the Comment object to send to repo
	return s.r.Create(ctx, c);
}

func (s *commentSvc) UpdateComment(ctx context.Context, c *Comment) error {
	return s.r.UpdateByID(ctx, c)
}

func (s *commentSvc) FindComments(ctx context.Context, postID uuid.UUID, page int, limit int) ([]*Comment, error) {
	return s.r.ReadByPostID(ctx, postID, page, limit)
}

func (s *commentSvc) FindByID(ctx context.Context, id uuid.UUID) (*Comment, error) {
	return s.r.ReadByID(ctx, id)
}

func (s *commentSvc) DeleteComment(ctx context.Context, id uuid.UUID) error {
	return s.r.DeleteByID(ctx, id)
}

func NewCommentSvc(r Repo) *commentSvc {
	return &commentSvc{r: r}
}