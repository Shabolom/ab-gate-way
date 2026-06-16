package di

import (
	"context"
	"fmt"

	"gate-way/internal/config"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type DI struct {
	config *config.Config
	logger *zap.Logger

	abConn   *grpc.ClientConn
	authConn *grpc.ClientConn

	grpcServer *grpc.Server

	//metrics *metrics.Metrics
}

func New(ctx context.Context) *DI {
	_ = ctx
	return &DI{}
}

func (d *DI) Config() *config.Config {
	if d.config != nil {
		return d.config
	}

	cfg, err := config.FromEnv()
	if err != nil {
		panic(fmt.Errorf("config from env: %w", err))
	}

	d.config = cfg
	return d.config
}

func (d *DI) Logger() *zap.Logger {
	if d.logger != nil {
		return d.logger
	}

	var logger *zap.Logger
	var err error

	if d.Config().Debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		panic(fmt.Errorf("create logger: %w", err))
	}

	logger = logger.With(
		zap.String("service", d.Config().ServiceName),
		zap.Bool("debug", d.Config().Debug),
	)

	d.logger = logger
	_ = zap.ReplaceGlobals(logger)

	return d.logger
}

func (d *DI) Shoutdown() {
	err := d.authConn.Close()

	if err != nil {
		d.Logger().Error("shoutdown", zap.Error(err))
	}

}
