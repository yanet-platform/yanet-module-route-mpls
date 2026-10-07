use core::error::Error;

fn main() -> Result<(), Box<dyn Error>> {
    // The include root is the yanet2 checkout this module sits in, so the
    // proto's common/commonpb imports resolve; the proto itself is this
    // repository's, at its submodule path. Standalone use adjusts both
    // (see README.md).
    ync_build::client(
        "../../..",
        &["external/route-mpls/controlplane/routemplspb/v1/routempls.proto"],
    )
    .serialize()
    .with(|builder| builder.enum_attribute(".", "#[derive(serde::Serialize)]"))
    .compile()
}
