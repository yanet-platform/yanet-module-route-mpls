package main

import (
	"go.uber.org/zap/zapcore"

	route_mpls "github.com/yanet-platform/yanet-module-route-mpls/controlplane"
	"github.com/yanet-platform/yanet2/common/go/logging"
	"github.com/yanet-platform/yanet2/common/go/operator"
	"github.com/yanet-platform/yanet2/common/go/xcfg"
)

// Config is the daemon's YAML configuration.
type Config struct {
	// Module carries the shared-memory attachment parameters passed to
	// ffi.Attach.
	Module route_mpls.Config `yaml:"module"`

	// Server is where the daemon serves the module's gRPC service.
	Server operator.GRPCServerConfig `yaml:"server"`

	// Register paces the gateway registration heartbeat.
	Register operator.RegisterConfig `yaml:"register"`

	// Gateways receive the registration heartbeat advertising Server.
	// An empty list leaves the daemon serving its endpoint directly,
	// reachable without the gateway.
	Gateways []operator.GatewayConfig `yaml:"gateways"`

	Logging logging.Config `yaml:"logging"`
}

// DefaultConfig returns default configuration.
func DefaultConfig() *Config {
	return &Config{
		Module: *route_mpls.DefaultConfig(),
		Server: operator.GRPCServerConfig{
			Endpoint: xcfg.MustNonEmptyString("[::1]:8091"),
		},
		Register: operator.RegisterConfig{
			Interval: xcfg.MustNonZero(operator.DefaultRegisterInterval),
		},
		Gateways: []operator.GatewayConfig{},
		Logging: logging.Config{
			Level: zapcore.InfoLevel,
		},
	}
}

// Default resets Config to DefaultConfig.
func (m *Config) Default() {
	*m = *DefaultConfig()
}

// LoggingConfig exposes the logging configuration to the generic operator
// CLI helper.
func (m *Config) LoggingConfig() *logging.Config {
	return &m.Logging
}
