package domain

import (
	"github.com/google/uuid"
)

type User struct {
	Base         `tstype:",extends"`
	TenantScoped `tstype:",extends"`
	Email        string    `json:"email"`
	RoleID       uuid.UUID `json:"roleID"`
	Type         UserType  `json:"type"`
}
