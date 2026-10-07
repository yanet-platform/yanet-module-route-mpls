# Control-plane FFI archive consumed by the Go bindings through their #cgo
# -L/-l flags. The dataplane config struct (dataplane/config.h) is embedded
# in the shm config, hence the include.
#
# The archive lands in its own build/api/ directory so that the -L path the
# Go bindings use never contains a shared object: a libroute_mpls_dp.so
# there would shadow the in-tree static module when the in-process test
# harness links -lroute_mpls_dp.
static_library(
  'route_mpls_cp',
  files('controlplane.c'),
  dependencies: [yanet_sdk],
  include_directories: include_directories('../dataplane'),
  install: false,
)
