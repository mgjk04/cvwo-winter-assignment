package post

import (
	"time"
	"github.com/google/uuid"
)

type Post struct {
	ID uuid.UUID `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	TopicID uuid.UUID `json:"topic_id"`
	AuthorID uuid.UUID `json:"author_id"`
	AuthorName string `json:"authorname"` 
}

type SearchQuery struct {
	Page int `form:"page,default=1" binding:"gt=0"`
	Limit int `form:"limit,default=10" binding:"gt=0"`
}

//DTOs
type PostCreateReq struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
	AuthorID uuid.UUID `json:"author_id" binding:"omitempty"`
	//topic_id not required as it is part of the route params
}
type PostUpdateReq struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
	TopicID uuid.UUID `json:"topic_id" binding:"omitempty"`
	AuthorID uuid.UUID `json:"author_id" binding:"omitempty"`
}

type PostReadRes struct {
	Posts []*Post `json:"posts"`
	Count int `json:"count"`
}