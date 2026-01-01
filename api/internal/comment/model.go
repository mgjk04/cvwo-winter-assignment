package comment

import (
	"time"
	"github.com/google/uuid"
)

type Comment struct {
	ID uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	PostID    uuid.UUID `json:"post_id"`
	AuthorID  uuid.UUID `json:"author_id"`
}

//NOTE: consider moving SearchQuery into its own file, repeated quite a bit 
type SearchQuery struct {
	Page int `form:"page,default=1" binding:"gt=0"`
	Limit int `form:"limit,default=10" binding:"gt=0"`
}

//DTOs
type CommentCreateReq struct {
	Content string `json:"content" binding:"required"`
	AuthorID uuid.UUID `json:"author_id" binding:"required"`
	//post_id not required as it is part of the route params
}
type CommentUpdateReq struct {
	Content string `json:"content" binding:"required"`
	PostID uuid.UUID `json:"post_id" binding:"required"`
	AuthorID uuid.UUID `json:"author_id" binding:"required"`
}