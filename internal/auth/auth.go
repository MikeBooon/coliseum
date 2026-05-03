package auth

import "github.com/MikeBooon/coliseum/internal/config"

type AuthClient struct {
	EnvConfig *config.EnvConfig
}

func NewAuthClient(env *config.EnvConfig) *AuthClient {
	return &AuthClient{
		EnvConfig: env,
	}
}
