# route-mpls, ported out of the yanet2 tree to the out-of-tree module shape.
#
# This is a standalone meson project: the yanet2 tree never traverses it.
# It builds against a configured yanet2 checkout through the module SDK
# (yanet-module-sdk.pc — point PKG_CONFIG_PATH at the checkout's build/sdk,
# which this repository's Makefile does). See README.md.
project(
  'yanet-module-route-mpls',
  'c',
  meson_version: '>= 0.61',
  default_options: [
    'buildtype=release',
    'warning_level=2',
    'werror=true',
    'optimization=2',
  ],
)

yanet_sdk = dependency('yanet-module-sdk')

# The plugin the dataplane dlopen's from plugin_dir as libroute_mpls_dp.so.
# Undefined references to yanet and DPDK symbols stay unresolved here and
# are satisfied by the dataplane binary's exported dynamic symbols at load
# time, exactly as for the statically linked built-ins.
#
# The .so keeps the build root to itself on purpose: the control-plane
# archive lives in build/api/ (see api/meson.build), so the CGO -L flag
# never exposes this directory's .so to the linker — the in-process test
# harness links the in-tree static route_mpls_dp, and a same-named .so on
# its search path would bind the test binary to the plugin instead.
shared_module(
  'route_mpls_dp',
  files('dataplane/dataplane.c'),
  dependencies: [yanet_sdk],
  install: false,
)

subdir('api')
