pub mod operator_requester;

use alloy::primitives::B256;
use anyhow::{anyhow, Result};
use ark_serialize::CanonicalDeserialize;
use ark_serialize::{Compress, Validate};
use eigensdk::client_avsregistry::reader::{AvsRegistryChainReader, AvsRegistryReader};
use eigensdk::crypto_bls::{BlsSignature, Signature};
use eigensdk::logging::get_logger;
use eigensdk::services_avsregistry::AvsRegistryService;
use eigensdk::services_blsaggregation::bls_agg::{
    BlsAggregatorService, TaskMetadata, TaskSignature,
};
use eigensdk::services_blsaggregation::bls_aggregation_service_response::BlsAggregationServiceResponse;
use eigensdk::types::avs::TaskIndex;
use eigensdk::types::operator::{QuorumNum, QuorumThresholdPercentage};
use ethers::utils::keccak256;
use operator_requester::OperatorRequester;
use std::io::Cursor;
use std::sync::{Arc, Mutex};
use std::time::Duration;

pub struct AggregatorService<
    T: AvsRegistryService + Send + Sync + Clone + 'static,
    R: OperatorRequester + Clone,
> {
    avs_chain_reader: AvsRegistryChainReader,
    avs_registry_service: T,
    operator_requester: R,
    mutex: Mutex<()>,
}

impl<
        T: AvsRegistryService + Send + Sync + Clone + 'static,
        R: OperatorRequester + Clone + 'static,
    > AggregatorService<T, R>
{
    pub fn new(
        avs_chain_reader: AvsRegistryChainReader,
        avs_registry_service: T,
        operator_requester: R,
    ) -> Self {
        Self {
            avs_chain_reader,
            avs_registry_service,
            operator_requester,
            mutex: Mutex::new(()),
        }
    }

    pub async fn get_certificate(
        &self,
        task_index: TaskIndex,
        task_created_block: u64,
        quorum_number: QuorumNum,
        quorum_threshold_percentage: QuorumThresholdPercentage,
        data: Vec<u8>,
        time_to_expiry: Duration,
    ) -> Result<BlsAggregationServiceResponse> {
        // Only allow one task at a time
        let _lock = self
            .mutex
            .lock()
            .map_err(|_| anyhow!("Failed to acquire lock"))?;

        let quorum_numbers = vec![quorum_number];
        let quorum_threshold_percentages = vec![quorum_threshold_percentage];

        let task_metadata = TaskMetadata::new(
            task_index.clone(),
            task_created_block,
            quorum_numbers.clone(),
            quorum_threshold_percentages,
            time_to_expiry,
        );

        // Initialize task in BLS aggregation service
        let bls_agg_service =
            BlsAggregatorService::new(self.avs_registry_service.clone(), get_logger());
        let (service_handle, mut aggregate_receiver) = bls_agg_service.start();
        service_handle.initialize_task(task_metadata).await?;
        let service_handle = Arc::new(service_handle);

        // Get operators from registry
        let operators = self
            .avs_registry_service
            .get_operators_avs_state_at_block(task_created_block, &quorum_numbers)
            .await?;

        tracing::info!(
            num_operators = operators.len(),
            "Number of operators registered"
        );

        // Send task to all operators in parallel
        let mut handles = Vec::new();
        for (operator_id, operator) in operators {
            let operator_address = self
                .avs_chain_reader
                .get_operator_from_id(*operator_id)
                .await?;
            let operator_requester = self.operator_requester.clone();
            let data = data.clone();
            let service_handle = service_handle.clone();

            let handle = tokio::spawn(async move {
                tracing::info!(operator_id = ?operator_id, "Requesting certification from operator");

                // Create connection for this operator
                let resp = match operator_requester
                    .request_certification(operator, operator_address, task_index, &data)
                    .await
                {
                    Ok(resp) => resp,
                    Err(_) => return Ok::<(), anyhow::Error>(()),
                };

                tracing::debug!(operator_id = ?operator_id, "Received signature from operator");

                let cursor = Cursor::new(&resp.signature);
                let signature: BlsSignature =
                    BlsSignature::deserialize_with_mode(cursor, Compress::No, Validate::Yes)
                        .map_err(|e| anyhow!("Failed to deserialize signature: {}", e))?;
                let signature = Signature::new(signature);

                let task_data = resp.data;
                let task_response_digest = B256::from_slice(&keccak256(&task_data));
                tracing::debug!(operator_id = ?operator_id, "Processing signature from operator");

                return match service_handle
                    .process_signature(TaskSignature::new(
                        task_index.clone(),
                        task_response_digest,
                        signature,
                        operator_id.clone(),
                    ))
                    .await
                {
                    Ok(_) => {
                        tracing::info!(operator_id = ?operator_id, "Processed signature from operator");
                        Ok::<(), anyhow::Error>(())
                    }
                    Err(err) => {
                        tracing::error!(operator_id = ?operator_id, error = ?err, "Failed to process signature");
                        return Ok::<(), anyhow::Error>(());
                    }
                };
            });

            handles.push(handle);
        }

        // Ensure all operator tasks complete (and log any panics)
        for handle in handles {
            if let Err(join_err) = handle.await {
                tracing::error!(error = ?join_err, "Operator task failed");
            }
        }

        // Await the aggregated response
        let result = aggregate_receiver.receive_aggregated_response().await?;
        tracing::info!("Received aggregated response");
        Ok(result)
    }
}
