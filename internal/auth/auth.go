package auth

import (
	"github.com/MikeBooon/coliseum/internal/config"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

type AuthClient struct {
	client *casdoorsdk.Client
}

func NewAuthClient(env *config.EnvConfig) *AuthClient {
	casdoor := casdoorsdk.NewClient(
		env.CasdoorDomain,
		env.CasdoorAppID,
		env.CasdoorAppSecret,
		env.CasdoorAppCert,
		env.CasdoorAdminOrg,
		env.CasdoorAppName,
	)

	return &AuthClient{
		client: casdoor,
	}
}
