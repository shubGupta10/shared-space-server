package models

import (
	"time"

	"github.com/google/uuid"
)

type Space struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Token     string    `json:"token" unique:"true"`
	Creator   uuid.UUID `json:"creator"`
	Partner   uuid.UUID `json:"partner"`
	CreatedAt time.Time `json:"created_at"`
}
