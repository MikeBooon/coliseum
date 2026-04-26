package dao

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	BaseColumns
	bun.BaseModel
	Email    string    `bun:"email,notnull"`
	TenantID uuid.UUID `bun:"tenant_id,notnull,type:uuid"`
}
