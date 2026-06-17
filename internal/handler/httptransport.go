package httptransport

import (
	customParams "gate-way/internal/handler/ab-microservice/custom-params"
	"gate-way/internal/handler/ab-microservice/experiments"
	"gate-way/internal/handler/ab-microservice/layer"
	"gate-way/internal/handler/ab-microservice/namespace"
	"gate-way/internal/handler/auth-microservice/auth"
	userHandler "gate-way/internal/handler/auth-microservice/user"
)

type (
	getAuthHandlers        = *auth.Handler
	getUserHandlers        = *userHandler.Handler
	getCustomParamHandlers = *customParams.Handler
	getExperimentHandlers  = *experiments.Handler
	getLayerHandlers       = *layer.Handler
	getNamespaceHandlers   = *namespace.Handler
)

type Handlers struct {
	getAuthHandlers
	getUserHandlers
	getCustomParamHandlers
	getExperimentHandlers
	getLayerHandlers
	getNamespaceHandlers
}

func New(
	getAuthHandlers getAuthHandlers,
	getUserHandlers getUserHandlers,
	getCustomParamHandlers getCustomParamHandlers,
	getExperimentHandlers getExperimentHandlers,
	getLayerHandlers getLayerHandlers,
	getNamespaceHandlers getNamespaceHandlers,
) *Handlers {
	return &Handlers{
		getAuthHandlers:        getAuthHandlers,
		getUserHandlers:        getUserHandlers,
		getCustomParamHandlers: getCustomParamHandlers,
		getExperimentHandlers:  getExperimentHandlers,
		getLayerHandlers:       getLayerHandlers,
		getNamespaceHandlers:   getNamespaceHandlers,
	}
}
