package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type MagickNumbers struct {
	OperatingInterval int `envconfig:"OPERATING_INTERVAL" default:"5"`
	ContextInterval   int `envconfig:"CONTEXT_INTERVAL" default:"30"`
}

type Prometheus struct {
	Port string `envconfig:"PROMETHEUS_PORT" default:"2112"`
}

type Config struct {
	ServiceName        string `envconfig:"APP_NAME"`
	Debug              bool   `envconfig:"APP_DEBUG"`
	GRPCPort           string `envconfig:"APP_GRPC_ADDRESS"`
	Secret             string `envconfig:"APP_SECRET"`
	ResendAppKey       string `envconfig:"RESEND_API_KEY"`
	HealthcheckPort    string `envconfig:"HEALTHCHECK_PORT" default:"8093"`
	Port               string `envconfig:"APP_PORT"`
	ABGrpcConnection   string `envconfig:"GRPC_CONNECTION"`
	AuthConnectionGRPC string `envconfig:"AUTH_CONNECTION_GRPC"`
	Prometheus         Prometheus
	MagickNumbers      MagickNumbers
}

func FromEnv() (*Config, error) {
	cfg := new(Config)

	if err := envconfig.Process("", cfg); err != nil {
		return nil, fmt.Errorf("error while parse env config | %w", err)
	}

	return cfg, nil
}
