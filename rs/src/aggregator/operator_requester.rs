use crate::api::proto::node::node_service_client::NodeServiceClient;
use crate::api::proto::node::{CertifyRequest, CertifyResponse};
use alloy::primitives::Address;
use anyhow::{anyhow, Result};
use async_trait::async_trait;
use eigensdk::types::avs::TaskIndex;
use eigensdk::types::avs_state::OperatorAvsState;
use tonic::Request;

#[async_trait]
pub trait OperatorRequester: Send + Sync {
    async fn request_certification(
        &self,
        operator: OperatorAvsState,
        operator_address: Address,
        task_index: TaskIndex,
        data: &[u8],
    ) -> Result<CertifyResponse>;
}

#[derive(Debug, Clone)]
pub struct GrpcOperatorRequester {}

impl GrpcOperatorRequester {
    pub fn new() -> Self {
        Self {}
    }
}

#[async_trait]
impl OperatorRequester for GrpcOperatorRequester {
    async fn request_certification(
        &self,
        operator: OperatorAvsState,
        operator_address: Address,
        task_index: TaskIndex,
        data: &[u8],
    ) -> Result<CertifyResponse> {
        let url = operator_address.to_string();
        let mut client = NodeServiceClient::connect(url)
            .await
            .map_err(|e| anyhow!("Failed to connect to operator: {}", e))?;
        let pb_req = CertifyRequest {
            task_index: task_index.into(),
            data: data.to_vec(),
        };
        let resp = client
            .certify(Request::new(pb_req))
            .await
            .map_err(|e| anyhow!("gRPC error: {}", e))?
            .into_inner();
        Ok(CertifyResponse {
            signature: resp.signature.to_vec(),
            data: resp.data.to_vec(),
        })
    }
}

// A dummy operator requester for testing
#[cfg(test)]
pub struct DummyOperatorRequester(pub CertifyResponse);

#[async_trait]
#[cfg(test)]
impl OperatorRequester for DummyOperatorRequester {
    async fn request_certification(
        &self,
        _operator: OperatorAvsState,
        _operator_address: Address,
        _task_index: TaskIndex,
        _data: &[u8],
    ) -> Result<CertifyResponse> {
        Ok(self.0.clone())
    }
}
