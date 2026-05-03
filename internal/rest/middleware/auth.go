package middleware

import (
	"context"

	"github.com/MikeBooon/coliseum/internal/system"
	"github.com/zitadel/zitadel-go/v3/pkg/authentication"
	"github.com/zitadel/zitadel-go/v3/pkg/authentication/oidc"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
)

const REDIRECT_PATH = "/api/v1/redirect"

func GetAuthMiddleware(ctx context.Context, sys *system.System) (*authentication.Interceptor[*oidc.DefaultContext], error) {
	auth, err := authentication.New(
		ctx,
		zitadel.New(sys.EnvConfig.ZitadelDomain),
		sys.EnvConfig.EncKey,
		oidc.DefaultAuthentication(sys.EnvConfig.ZitadelClient, sys.EnvConfig.Domain+REDIRECT_PATH, sys.EnvConfig.EncKey),
	)

	if err != nil {
		return nil, err
	}

	return authentication.Middleware(auth), nil
}
