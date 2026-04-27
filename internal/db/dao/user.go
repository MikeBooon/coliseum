package dao

import (
	"github.com/MikeBooon/coliseum/domain"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	Base
	bun.BaseModel
	TenantScoped
	Email  string          `bun:"email,notnull"`
	RoleID uuid.UUID       `bun:"role_id,notnull,type:uuid"`
	Type   domain.UserType `bun:"type,notnull,type:user_type"`
}

func (u *User) ToDomain() *domain.User {
	return &domain.User{
		Base:         u.Base.ToDomain(),
		TenantScoped: u.TenantScoped.ToDomain(),
		Email:        u.Email,
		Type:         u.Type,
		RoleID:       u.RoleID,
	}
}
