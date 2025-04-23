use anyhow::Result;
use ark_serialize::CanonicalSerialize;
use ethers::utils::keccak256;
use serde::Serialize;
use tokio::signal;
use tonic::{Request, Response, Status};

use eigensdk::crypto_bls::BlsKeyPair;

use crate::api::proto::node::{node_service_server::NodeService, CertifyRequest, CertifyResponse};

pub struct CertifyingService {
    key_pair: BlsKeyPair,
    get_response: Box<dyn Fn(&[u8]) -> Result<Vec<u8>> + Send + Sync>,
}

impl CertifyingService {
    pub fn new(
        key_pair: BlsKeyPair,
        get_response: Box<dyn Fn(&[u8]) -> Result<Vec<u8>> + Send + Sync>,
    ) -> Self {
        Self {
            key_pair,
            get_response,
        }
    }
}

#[tonic::async_trait]
impl NodeService for CertifyingService {
    async fn certify(
        &self,
        request: Request<CertifyRequest>,
    ) -> Result<Response<CertifyResponse>, Status> {
        let data = request.into_inner().data;

        let response = (self.get_response)(&data)
            .map_err(|e| Status::invalid_argument(format!("Data is invalid: {}", e)))?;

        let digest = keccak256(&response);
        let signature = self.key_pair.sign_message(&digest);

        // Serialize to raw (uncompressed) bytes, like the go version.
        let mut signature_bytes = Vec::new();
        signature
            .g1_point()
            .g1()
            .serialize_uncompressed(&mut signature_bytes)
            .map_err(|e| Status::internal(format!("serialization error: {}", e)))?;

        Ok(Response::new(CertifyResponse {
            signature: signature_bytes,
            data: response,
        }))
    }
}
