# yanet-module-route-mpls

The [YANET](https://github.com/yanet-platform/yanet2) route-mpls module,
ported to the out-of-tree module shape: the dataplane is a runtime plugin
(`libroute_mpls_dp.so`), the control plane is a standalone daemon that
registers itself with the gateway, and the CLI is a standalone cargo
workspace. It builds against a yanet2 checkout that carries the [module
SDK](https://github.com/yanet-platform/yanet2/blob/main/docs/module-sdk.md)
— nothing in the yanet2 tree needs editing.

This repository is the production-grade sibling of yanet2's
`sdk/example`: a real module with filters, tunnel nexthops and a full
functional suite, built strictly through the SDK surfaces.

## Layout

```
api/                  C library the control-plane FFI calls (unchanged from
                      the in-tree module)
bindings/go/croutempls CGO wrapper over api/
controlplane/          Go control plane: gRPC service, shm write path, protos
cmd/route-mpls-controlplane  the out-of-tree control-plane daemon
dataplane/             dataplane plugin sources (unchanged from the in-tree
                      module)
cli/                   yanet-cli-route-mpls, standalone cargo workspace
tests/functional/      in-process dataplane tests, plugin-loaded
etc/                   sample daemon configuration
```

## Building

The module's default home is the `external/route-mpls` submodule of a
yanet2 checkout, and every relative path (the go.mod `replace`, the CLI's
`ync` path dependency, the CGO include paths) assumes that location.

```sh
# once, in the yanet2 checkout (needs the module SDK):
git submodule update --init && meson setup build && meson compile -C build

# in this repository:
make test       # plugin + FFI archive + daemon + CLI + all tests
```

Building from a different location: point `YANET_ROOT` at the checkout,
adjust the `replace` directive in `go.mod`, the paths in
`bindings/go/croutempls/cgo.go` and `cli/Cargo.toml`, and export
`PKG_CONFIG_PATH=$YANET_ROOT/build/sdk`.

## Deploying

- Copy `build/libroute_mpls_dp.so` into the dataplane's `plugin_dir` and
  name `route_mpls` in the dataplane's `modules:` list.
- Run `route-mpls-controlplane --config etc/route-mpls-controlplane.yaml`;
  its gateway entry makes the module's service reachable through the
  gateway exactly like an in-tree module's.

## Status

Experimental port. The in-tree module (yanet2's `modules/route-mpls`)
remains the shipped implementation; this repository exists to prove and
pressure-test the out-of-tree module SDK with a production module.
