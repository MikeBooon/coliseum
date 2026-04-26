package dao

import (
	"time"

	"github.com/MikeBooon/coliseum/domain"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Role struct {
	Base
	bun.BaseModel
	Name     string          `bun:"name,notnull"`
	TenantID uuid.UUID       `bun:"tenant_id,notnull,type:uuid"`
	Type     domain.UserType `bun:"type,notnull,type:user_type"`
}

type Permission struct {
	bun.BaseModel
	ID        uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	Key       string    `bun:"key,notnull"`
	RoleID    uuid.UUID `bun:"role_id,notnull,type:uuid"`
}
