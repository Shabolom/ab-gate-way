package config

import (
	"fmt"
	"net"

	"github.com/kelseyhightower/envconfig"
)

type MagickNumbers struct {
	OperatingInterval int `envconfig:"OPERATING_INTERVAL" default:"5"`
	ContextInterval   int `envconfig:"CONTEXT_INTERVAL" default:"30"`
}

type Prometheus struct {
	Port string `envconfig:"PROMETHEUS_PORT" default:"2112"`
}

type ABConnectionGRPC struct {
	Host string `envconfig:"AB_GRPC_HOST" default:"localhost"`
	Port string `envconfig:"AB_GRPC_PORT" default:"8015"`
}

type AuthConnectionGRPC struct {
	Host string `envconfig:"AUTH_GRPC_HOST" default:"localhost"`
	Port string `envconfig:"AUTH_GRPC_PORT" default:"8019"`
}

type Config struct {
	ServiceName        string `envconfig:"APP_NAME"`
	Debug              bool   `envconfig:"APP_DEBUG"`
	GRPCPort           string `envconfig:"APP_GRPC_ADDRESS"`
	Secret             string `envconfig:"APP_SECRET"`
	ResendAppKey       string `envconfig:"RESEND_API_KEY"`
	HealthcheckPort    string `envconfig:"HEALTHCHECK_PORT" default:"8093"`
	Port               string `envconfig:"APP_PORT"`
	GrpcConnection     ABConnectionGRPC
	AuthConnectionGRPC AuthConnectionGRPC
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

func (c *Config) GrpcAddressAB() string {
	return net.JoinHostPort(
		c.GrpcConnection.Host,
		c.GrpcConnection.Port,
	)
}

func (c *Config) GrpcAddressAuth() string {
	return net.JoinHostPort(
		c.GrpcConnection.Host,
		c.GrpcConnection.Port,
	)
}
