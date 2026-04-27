package tenant

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type tenantContextKey struct{}

var key = tenantContextKey{}

var ErrTenantNotInContext = errors.New("tenant not in context")

func NewContext(ctx context.Context, tenantID uuid.UUID) context.Context {
	return context.WithValue(ctx, key, tenantID)
}

func FromContext(ctx context.Context) (uuid.UUID, error) {
	id, ok := ctx.Value(key).(uuid.UUID)
	if !ok {
		return uuid.UUID{}, ErrTenantNotInContext
	}
	return id, nil
}

func MustFromContext(ctx context.Context) uuid.UUID {
	id, ok := ctx.Value(key).(uuid.UUID)
	if !ok {
		panic("tenant missing from context")
	}
	return id
}
