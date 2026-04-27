package controller

import (
	"github.com/MikeBooon/coliseum/domain/dto"
	"github.com/MikeBooon/coliseum/internal/rest/input"
	"github.com/MikeBooon/coliseum/internal/system"
	"github.com/MikeBooon/coliseum/service"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type AuthController struct {
	s *service.Services
}

const RATE_LIMIT = 4

var authRateLimiter = middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(RATE_LIMIT))

func NewAuthController(g *echo.Group, s system.System) *AuthController {
	ctrl := &AuthController{s: s.Svcs}

	group := g.Group("")
	group.Use(authRateLimiter)

	g.POST("/auth/login", ctrl.Login)

	return ctrl
}

func (ctrl *AuthController) Login(c *echo.Context) error {
	body := new(dto.LoginDTO)
	if err := input.BindAndValidate(c, body); err != nil {
		return err
	}
	return nil
}
