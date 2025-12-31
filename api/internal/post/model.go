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
}