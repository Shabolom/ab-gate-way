package di

import (
	httptransport "gate-way/internal/handler"
	customParams "gate-way/internal/handler/ab-microservice/custom-params"
	"gate-way/internal/handler/ab-microservice/experiments"
	"gate-way/internal/handler/ab-microservice/layer"
	"gate-way/internal/handler/ab-microservice/namespace"
	"gate-way/internal/handler/auth-microservice/auth"
	"gate-way/internal/handler/auth-microservice/user"
)

func (d *DI) GetHandlersHTTP() *httptransport.Handlers {
	return httptransport.New(
		d.getAuthHandlers(),
		d.getUserHandlers(),
		d.getCustomParamHandlers(),
		d.getExperimentHandlers(),
		d.getLayerHandlers(),
		d.getNamespaceHandlers(),
	)
}

func (d *DI) getUserHandlers() *user.Handler {
	return user.New(d.GetAuthService())
}

func (d *DI) getAuthHandlers() *auth.Handler {
	return auth.New(d.GetAuthService())
}

func (d *DI) getLayerHandlers() *layer.Handler {
	return layer.New(d.GetAbService())
}

func (d *DI) getNamespaceHandlers() *namespace.Handler {
	return namespace.New(d.GetAbService())
}

func (d *DI) getExperimentHandlers() *experiments.Handler {
	return experiments.New(d.GetAbService())
}

func (d *DI) getCustomParamHandlers() *customParams.Handler {
	return customParams.New(d.GetAbService())
}
