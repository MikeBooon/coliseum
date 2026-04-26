package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Base     `tstype:",extends"`
	Email    string    `json:"email"`
	TenantID uuid.UUID `json:"tenantID"`
	RoleID   uuid.UUID `json:"roleID"`
	Type     UserType  `json:"type"`
}
