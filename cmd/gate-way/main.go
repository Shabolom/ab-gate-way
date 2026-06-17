package main

import (
	"context"
	api "gate-way/gen/echo"
	"gate-way/internal/di"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	if err := godotenv.Load("./build/local/.env"); err != nil {
		panic(err)
	}

	e := echo.New()

	container := di.New(ctx)
	container.Logger()

	handlers := container.GetHandlersHTTP()

	api.RegisterHandlers(e, handlers)

	go func() {
		err := e.Start(":" + container.Config().Port)
		if err != nil {
			container.Logger().Fatal("shutting down the server", zap.Error(err))
		}
	}()

	<-ctx.Done()
	container.Logger().Info("shutting down the server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(),
		time.Duration(container.Config().MagickNumbers.ContextInterval)*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		container.Logger().Error("failed to shutdown http server", zap.Error(err))
	}

	container.Shutdown()

	container.Logger().Info("shutting down gracefully")
}
