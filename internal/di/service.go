package di

import (
	"gate-way/internal/service/ab"
	"gate-way/internal/service/auth"
)

func (d *DI) GetAuthService() *auth.Service {
	return auth.New(d.GetAuthAdapter(), d.Logger())
}

func (d *DI) GetAbService() *ab.Service {
	return ab.New(d.GetAbAdapter(), d.Logger())
}
