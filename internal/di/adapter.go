package di

import (
	abAdapter "gate-way/internal/adapter/ab"
	authAdapter "gate-way/internal/adapter/auth"
)

func (d *DI) GetAuthAdapter() *authAdapter.Adapter {
	return authAdapter.New(d.NewAuthClientGRPC())
}

func (d *DI) GetAbAdapter() *abAdapter.Adapter {
	return abAdapter.New(d.NewAbClientGRPC())
}
