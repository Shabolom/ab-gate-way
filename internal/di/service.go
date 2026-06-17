package di

import (
	"gate-way/internal/service/ab"
	"gate-way/internal/service/auth"
)

func (d *DI) GetAuthService() *authService.Service {
	return authService.New(d.GetAuthAdapter(), d.logger)
}

func (d *DI) GetAbService() *ab.Service {
	return ab.New(d.GetAbAdapter(), d.logger)
}
