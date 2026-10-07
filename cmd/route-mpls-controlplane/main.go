// Command route-mpls-controlplane is the control-plane daemon of the
// out-of-tree route-mpls module.
//
// The module is not compiled into yncp-director: this daemon attaches to
// the dataplane's shared memory, serves the module's gRPC service on its
// own endpoint, and heartbeats itself into the gateway's backend registry
// so CLIs routed through the gateway reach it. See README.md.
package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	_ "google.golang.org/grpc/encoding/gzip"

	route_mpls "github.com/yanet-platform/yanet-module-route-mpls/controlplane"
	"github.com/yanet-platform/yanet2/common/go/operator"
)

func main() {
	err := operator.Run[Config](
		"yanet-route-mpls-controlplane",
		"YANET route-mpls module control plane (out-of-tree port)",
		buildDaemon,
	)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}

func buildDaemon(cfg *Config, log *zap.Logger) (operator.Runnable, error) {
	module, err := route_mpls.NewRouteMPLSModule(&cfg.Module, route_mpls.WithLog(log))
	if err != nil {
		return nil, err
	}

	return newDaemon(cfg, module, WithLog(log))
}
