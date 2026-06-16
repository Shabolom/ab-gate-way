package handler

import (
	"gate-way/internal/handler/auth-microservice/auth"
	userHandler "gate-way/internal/handler/auth-microservice/user"
)

type (
	getAuthHandlers = *auth.Handler
	getUserHandlers = *userHandler.Handler
)

type Handlers struct {
	getAuthHandlers
	getUserHandlers
}

func New(
	getAuthHandlers getAuthHandlers,
	getUserHandlers getUserHandlers,
) *Handlers {
	return &Handlers{
		getAuthHandlers: getAuthHandlers,
		getUserHandlers: getUserHandlers,
	}
}
