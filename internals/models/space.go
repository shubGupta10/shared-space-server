package models

import (
	"time"

	"github.com/google/uuid"
)

type Space struct {
	ID         uuid.UUID `json:"id"`
	Token      string    `json:"token" unique:"true"`
	CreatorOne uuid.UUID `json:"creator_one"`
	CreatorTwo uuid.UUID `json:"creator_two"`
	CreatedAt  time.Time `json:"created_at"`
}
