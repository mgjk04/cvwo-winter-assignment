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