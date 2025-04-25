use anyhow::Result;
use clap::Parser;
use eigensdk::crypto_bls::BlsKeyPair;
use teal::node::server::{BaseNode, Certifier, Config, Node};

#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]
struct Args {
    /// Port to listen on
    #[clap(short, long, default_value_t = 9000)]
    port: u16,
}

#[derive(Clone)]
struct SimpleCertifier;

impl Certifier for SimpleCertifier {
    fn get_response(&self, _config: &Config, data: &[u8]) -> Result<Vec<u8>> {
        // Simply echo the data back
        Ok(data.to_vec())
    }
}

#[tokio::main]
async fn main() -> Result<()> {
    // Initialize logging
    tracing_subscriber::fmt::init();

    // Parse command line arguments
    let args = Args::parse();

    // Create BLS key pair (use a hex string or your preferred seed)
    let key_pair =
        BlsKeyPair::new("0x1".to_string()).expect("Failed to create BLS key pair from seed");

    // Create node configuration
    let config = Config {
        service_port: args.port,
        bls_key_pair: key_pair,
    };

    // Create a simple certifier
    let certifier = SimpleCertifier;

    // Create and start the node
    let node = BaseNode::new(config, certifier);
    node.start().await?;

    Ok(())
}
