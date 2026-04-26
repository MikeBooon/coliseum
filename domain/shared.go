package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Base struct {
	ID        uuid.UUID    `json:"id"`
	CreatedAt uuid.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt bun.NullTime `bun:"deleted_at,soft_delete"`
}
