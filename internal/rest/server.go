// Package rest implements the rest server
package rest

import (
	"fmt"

	"github.com/labstack/echo/v5"
)

type Rest struct {
	e *echo.Echo
}

func NewRest() *Rest {
	e := echo.New()
	r := Rest{e}
	return &r
}

func (r *Rest) Start() error {
	if err := r.e.Start(":6464"); err != nil {
		return fmt.Errorf("eailed to start server: %w", err)
	}
	return nil
}
