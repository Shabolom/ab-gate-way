package di

import (
	"gate-way/internal/handler/auth-microservice/auth"
	"gate-way/internal/handler/auth-microservice/user"
)

func (d *DI) GetUserHandlers() *user.Handler {
	return user.New(d.GetAuthService())
}

func (d *DI) GetAuthHandlers() *auth.Handler {
	return auth.New(d.GetAuthService())
}
