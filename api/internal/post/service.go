package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/mgjk04/cvwo-winter-assignment/api/internal/generalErrors"
)

type Service interface {
	CreatePost(ctx context.Context, p *Post) (uuid.UUID, error)
	UpdatePost(ctx context.Context, userID uuid.UUID, p *Post) error
	FindPosts(ctx context.Context, topicID uuid.UUID, page int, limit int) ([]*Post, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Post, error)
	DeletePost(ctx context.Context, userID uuid.UUID, id uuid.UUID) error
}

type postSvc struct {
	r Repo 
}

func (s *postSvc) CreatePost(ctx context.Context, p *Post) (uuid.UUID, error) {
	return s.r.Create(ctx, p);
}

func (s *postSvc) UpdatePost(ctx context.Context, userID uuid.UUID, p *Post) error {
	isAuthor, err := s.verifyAuthor(ctx, userID, p.ID)
	if err != nil {
		return err
	}
	if !isAuthor {
		return generalErrors.ErrForbidden
	}
	return s.r.UpdateByID(ctx, p)
}

func (s *postSvc) FindPosts(ctx context.Context, topicID uuid.UUID, page int, limit int) ([]*Post, error) {
	return s.r.ReadByTopicID(ctx, topicID, page, limit)
}

func (s *postSvc) FindByID(ctx context.Context, id uuid.UUID) (*Post, error) {
	return s.r.ReadByID(ctx, id)
}

func (s *postSvc) DeletePost(ctx context.Context, userID uuid.UUID, postID uuid.UUID) error {
	isAuthor, err := s.verifyAuthor(ctx, userID, postID)
	if err != nil {
		return err
	}
	if !isAuthor {
		return generalErrors.ErrForbidden
	}
	return s.r.DeleteByID(ctx, postID)
}

func (s *postSvc) verifyAuthor(ctx context.Context, userID uuid.UUID, postID uuid.UUID) (bool, error) {
	post, err := s.r.ReadByID(ctx, postID)
	if err != nil {
		return false, err
	}
	return post.AuthorID == userID, err
}

func NewPostSvc(r Repo) *postSvc {
	return &postSvc{r: r}
}