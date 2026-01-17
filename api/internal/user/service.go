package user

import (
	"context"
	"github.com/google/uuid"
)

type Service interface {
	RegisterUser(ctx context.Context, username string) (uuid.UUID, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	DeRegisterUser(ctx context.Context, id uuid.UUID) error
}

type userSvc struct {
	r Repo 
}

func (s *userSvc) RegisterUser(ctx context.Context, username string) (uuid.UUID, error) {
	// this thing is in charge of assembbling the user object to send to repo
	return s.r.Create(ctx, &User{Username: username});
}

func (s *userSvc) FindByUsername(ctx context.Context, username string) (*User, error) {
	return s.r.ReadByUsername(ctx, username)
}

func (s *userSvc) FindByID(ctx context.Context, id uuid.UUID) (*User, error) {
	return s.r.ReadByID(ctx, id)
}

func (s *userSvc) DeRegisterUser(ctx context.Context, id uuid.UUID) error {
	return s.r.DeleteByID(ctx, id)
}

func NewUserSvc(r Repo) *userSvc {
	return &userSvc{r: r}
}