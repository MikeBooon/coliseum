package dao

import (
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
	Base
	bun.BaseModel
	Key    string    `bun:"key,notnull"`
	RoleID uuid.UUID `bun:"role_id,notnull,type:uuid"`
}
