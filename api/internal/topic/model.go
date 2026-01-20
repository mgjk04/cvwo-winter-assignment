package topic

import (
	"github.com/google/uuid"
	"time"
)

type Topic struct {
	ID uuid.UUID `json:"id"`
	TopicName string `json:"topicname"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	AuthorID uuid.UUID `json:"author_id"`
}

type SearchQuery struct {
	Page int `form:"page,default=1" binding:"gt=0"`
	Limit int `form:"limit,default=10" binding:"gt=0"`
}

//DTOs
type TopicCreateReq struct {
	TopicName string `json:"topicname" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
}
//ok thinking about it they're like the same, I'll keep it here for now
type TopicUpdateReq struct {
	TopicName string `json:"topicname" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
	AuthorID uuid.UUID `json:"author_id" binding:"omitempty"`
}

type TopicsReadRes struct {
	Topics []*Topic `json:"topics"`
	Count int `json:"count"`
}