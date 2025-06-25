package models

import (
	"time"

	"github.com/google/uuid"
)

type Space struct {
	ID        uuid.UUID `json:"id"`
	Token     string    `json:"token" unique:"true"`
	Creator   uuid.UUID `json:"creator"`
	Partner   uuid.UUID `json:"partner"`
	CreatedAt time.Time `json:"created_at"`
}
