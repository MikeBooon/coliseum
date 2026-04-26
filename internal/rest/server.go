// Package rest implements the rest server package rest
package rest

import (
	"fmt"

	"github.com/MikeBooon/coliseum/internal/rest/controller"
	"github.com/MikeBooon/coliseum/internal/system"
	"github.com/labstack/echo/v5"
)

type Rest struct {
	e *echo.Echo
}

func NewRest(sys system.System) *Rest {
	e := echo.New()

	baseGroupV1 := e.Group("/api/v1")
	controller.NewAuthController(baseGroupV1, sys)

	r := Rest{e}
	return &r
}

func (r *Rest) Start() error {
	if err := r.e.Start(":6464"); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}
