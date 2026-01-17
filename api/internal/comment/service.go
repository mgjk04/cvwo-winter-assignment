package comment

import (
	"context"

	"github.com/google/uuid"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/generalErrors"
)

type Service interface {
	CreateComment(ctx context.Context, p *Comment) (uuid.UUID, error)
	UpdateComment(ctx context.Context, userID uuid.UUID, p *Comment) error
	FindComments(ctx context.Context, topicID uuid.UUID, page int, limit int) ([]*Comment, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Comment, error)
	DeleteComment(ctx context.Context, userID uuid.UUID, commentID uuid.UUID) error
}

type commentSvc struct {
	r Repo 
}

func (s *commentSvc) CreateComment(ctx context.Context, c *Comment) (uuid.UUID, error) {
	// this thing is in charge of assembbling the Comment object to send to repo
	return s.r.Create(ctx, c);
}

func (s *commentSvc) UpdateComment(ctx context.Context, userID uuid.UUID, c *Comment) error {
	isAuthor, err := s.verifyAuthor(ctx, userID, c.ID)
	if err != nil {
		return err
	}
	if !isAuthor {
		return generalErrors.ErrForbidden
	}
	return s.r.UpdateByID(ctx, c)
}

func (s *commentSvc) FindComments(ctx context.Context, postID uuid.UUID, page int, limit int) ([]*Comment, error) {
	return s.r.ReadByPostID(ctx, postID, page, limit)
}

func (s *commentSvc) FindByID(ctx context.Context, id uuid.UUID) (*Comment, error) {
	return s.r.ReadByID(ctx, id)
}

func (s *commentSvc) DeleteComment(ctx context.Context,  userID uuid.UUID, commentID uuid.UUID) error {
	isAuthor, err := s.verifyAuthor(ctx, userID, commentID)
	if err != nil {
		return err
	}
	if !isAuthor {
		return generalErrors.ErrForbidden
	}
	return s.r.DeleteByID(ctx, commentID)
}

func (s *commentSvc) verifyAuthor(ctx context.Context, userID uuid.UUID, commentID uuid.UUID) (bool, error) {
	comment, err := s.r.ReadByID(ctx, commentID)
	if err != nil {
		return false, err
	}
	return comment.AuthorID == userID, err
}

func NewCommentSvc(r Repo) *commentSvc {
	return &commentSvc{r: r}
}