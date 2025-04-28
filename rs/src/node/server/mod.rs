use anyhow::{anyhow, Result};
use std::net::SocketAddr;
use std::sync::Arc;
use tokio::net::TcpListener;
use tonic::transport::Server;

use crate::{
    api::proto::node::node_service_server::NodeServiceServer, node::service::CertifyingService,
};
use eigensdk::crypto_bls::BlsKeyPair;

#[derive(Clone)]
pub struct Config {
    pub service_port: u16,
    pub bls_key_pair: BlsKeyPair,
}

pub trait Certifier: Clone + Send + Sync {
    fn get_response(&self, config: &Config, data: &[u8]) -> Result<Vec<u8>>;
}

pub struct BaseNode<C: Certifier> {
    config: Config,
    certifier: C,
}

#[async_trait::async_trait]
pub trait Node: Send + Sync {
    async fn start(&self) -> Result<()>;
}

impl<C: Certifier + 'static> BaseNode<C> {
    pub fn new(config: Config, certifier: C) -> Self {
        Self { config, certifier }
    }

    pub async fn start(&self) -> Result<()> {
        let addr = SocketAddr::from(([0, 0, 0, 0], self.config.service_port));
        self.start_with_listener(addr).await
    }

    pub async fn start_with_listener(&self, addr: SocketAddr) -> Result<()> {
        let certifier = self.certifier.clone();
        let config = self.config.clone();

        // Create a closure that captures the config for validation
        let get_response =
            move |data: &[u8]| -> Result<Vec<u8>> { certifier.get_response(&config, data) };

        let service =
            CertifyingService::new(self.config.bls_key_pair.clone(), Box::new(get_response));

        // Create a simplified router for our service
        // This is a simplified approach for the port example
        let router = tower::ServiceBuilder::new().service(service);

        tracing::info!(port = self.config.service_port, "Starting server");
        tracing::info!(addr = addr.to_string(), "Server listening on");

        // Use a simpler approach for the server
        tonic::transport::Server::builder()
            .add_service(NodeServiceServer::new(router))
            .serve(addr)
            .await
            .map_err(|e| anyhow!("Server error: {}", e))
    }
}
