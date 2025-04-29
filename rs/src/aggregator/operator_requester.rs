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
pub struct GrpcOperatorRequester {
    operator_url: String,
}

impl GrpcOperatorRequester {
    pub fn new(operator_url: String) -> Self {
        Self { operator_url }
    }
}

#[async_trait]
impl OperatorRequester for GrpcOperatorRequester {
    async fn request_certification(
        &self,
        // TODO pull the url from the OperatorInfo once we have it
        _operator: OperatorAvsState,
        operator_address: Address,
        task_index: TaskIndex,
        data: &[u8],
    ) -> Result<CertifyResponse> {
        tracing::info!(operator_address = ?operator_address, operator_url = %self.operator_url, "Connecting to operator");
        let mut client = NodeServiceClient::connect(self.operator_url.clone())
            .await
            .map_err(|e| anyhow!("Failed to connect to operator: {}", e))?;
        let pb_req = CertifyRequest {
            task_index: task_index.into(),
            data: data.to_vec(),
        };

        tracing::info!(pb_req = ?pb_req, "Sending request to operator");
        let resp = client
            .certify(Request::new(pb_req))
            .await
            .map_err(|e| anyhow!("gRPC error: {}", e))?
            .into_inner();
        tracing::info!(resp = ?resp, "Received response from operator");
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
