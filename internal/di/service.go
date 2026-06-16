package di

import "gate-way/internal/service/auth"

func (d *DI) GetAuthService() *authService.Service {
	return authService.New(d.GetAuthAdapter(), d.logger)
}
