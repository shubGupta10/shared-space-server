package models

import "github.com/google/uuid"

type Notes struct {
	ID        uuid.UUID `json:"id"`
	SpaceID   uuid.UUID `json:"space_id"`
	Content   string    `json:"content"`
	Author    uuid.UUID `json:"author"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}
