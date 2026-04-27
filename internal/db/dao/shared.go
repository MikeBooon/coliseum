package dao

import (
	"time"

	"github.com/MikeBooon/coliseum/domain"
	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	CreatedAt time.Time  `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time  `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
	DeletedAt *time.Time `bun:",soft_delete,nullzero"`
}

func (base Base) ToDomain() domain.Base {
	return domain.Base{
		ID:        base.ID,
		CreatedAt: base.CreatedAt,
		UpdatedAt: base.UpdatedAt,
		DeletedAt: base.DeletedAt,
	}
}

type TenantSetter interface {
	SetTenantID(id uuid.UUID)
}

type TenantScoped struct {
	TenantID uuid.UUID `bun:"tenant_id,notnull,type:uuid"`
}

func (base *TenantScoped) SetTenantID(tenantID uuid.UUID) {
	base.TenantID = tenantID
}

func (base TenantScoped) ToDomain() domain.TenantScoped {
	return domain.TenantScoped{
		TenantID: base.TenantID,
	}
}
