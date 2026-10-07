// The module path is this repository; the replace directive pins the yanet2
// packages to the checkout this module sits in (its default home is the
// external/route-mpls submodule of a yanet2 checkout, where ../../ is the
// checkout root). Standalone use points it at a local checkout instead.
module github.com/yanet-platform/yanet-module-route-mpls

go 1.27.1

replace github.com/yanet-platform/yanet2 => ../../

require (
	github.com/c2h5oh/datasize v0.0.0-20231215233829-aa82cc1e6500
	github.com/gopacket/gopacket v1.7.4
	github.com/stretchr/testify v1.12.1
	github.com/yanet-platform/xnetip v0.1.1
	github.com/yanet-platform/yanet2 v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.28.0
	golang.org/x/sync v0.23.0
	google.golang.org/grpc v1.84.0
	google.golang.org/protobuf v1.36.11
)

require (
	github.com/cenkalti/backoff/v5 v5.0.3 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/klauspost/compress v1.18.6 // indirect
	github.com/siderolabs/grpc-proxy v0.5.2 // indirect
	github.com/spf13/cobra v1.10.2 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	github.com/stripe/krl v0.0.0-20250403164848-fcf8101b2f53 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/crypto v0.54.0 // indirect
	golang.org/x/net v0.57.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/term v0.45.0 // indirect
	golang.org/x/text v0.40.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260706201446-f0a921348800 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
