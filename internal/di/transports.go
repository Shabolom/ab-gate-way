package di

import "gate-way/internal/handler"

func (d *DI) GetHTTPHandlers() *handler.HTTPHandlers {
	return handler.NewHTTPHandlers(
		d.GetExampleService(),
	)
}
