fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("cargo:rerun-if-changed=../proto/node.proto");

    tonic_build::configure()
        .build_server(true)
        .build_client(true)
        .compile(&["../proto/node.proto"], &["../proto"])
        .expect("Failed to compile node.proto");

    Ok(())
}
