//! End-to-end integration test for BLS aggregation

use alloy::network::EthereumWallet;
use alloy::primitives::{Address, Bytes, FixedBytes, U256};
use alloy::signers::local::PrivateKeySigner;
use anyhow::Result;
use eigensdk::client_avsregistry::reader::AvsRegistryChainReader;
use eigensdk::client_avsregistry::writer::AvsRegistryChainWriter;
use eigensdk::crypto_bls::BlsKeyPair;
use eigensdk::logging::{get_logger, get_test_logger};
use eigensdk::services_avsregistry::chaincaller::AvsRegistryServiceChainCaller;
use eigensdk::services_operatorsinfo::operatorsinfo_inmemory::OperatorInfoServiceInMemory;
use eigensdk::testing_utils::anvil_constants::{
    get_delegation_manager_address, get_erc20_mock_strategy, get_operator_state_retriever_address,
    get_registry_coordinator_address, get_rewards_coordinator_address, get_service_manager_address,
};
use eigensdk::testing_utils::transaction::wait_transaction;
use ethers::providers::{Http, Middleware, Provider};
use ethers::utils::Anvil;
use even_loving::EvenLovingCertifier;
use std::str::FromStr;
use std::time::{Duration, SystemTime, UNIX_EPOCH};
use teal::node::server::{BaseNode, Config};
use tokio::spawn;

pub mod even_loving;

use teal::{aggregator::operator_requester::GrpcOperatorRequester, AggregatorService};

const ANVIL_FIRST_PRIVATE_KEY: &str =
    "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80";

struct ContractAddresses {
    service_manager: Address,
    registry_coordinator: Address,
    operator_state_retriever: Address,
}

async fn get_contract_addresses(http_endpoint: String) -> Result<ContractAddresses> {
    let service_manager = get_service_manager_address(http_endpoint.clone()).await;
    let registry_coordinator = get_registry_coordinator_address(http_endpoint.clone()).await;
    let operator_state_retriever =
        get_operator_state_retriever_address(http_endpoint.clone()).await;

    Ok(ContractAddresses {
        service_manager,
        registry_coordinator,
        operator_state_retriever,
    })
}

async fn register_operator(
    avs_writer: &AvsRegistryChainWriter,
    bls_key_pair: BlsKeyPair,
    quorum_nums: Bytes,
    http_url: String,
) {
    let digest_hash: FixedBytes<32> = FixedBytes::from([0x02; 32]);

    // this is set to U256::MAX so that the registry does not take the signature as expired.
    let signature_expiry = U256::MAX;
    let tx_hash = avs_writer
        .register_operator_in_quorum_with_avs_registry_coordinator(
            bls_key_pair,
            digest_hash,
            signature_expiry,
            quorum_nums.clone(),
            "".into(),
        )
        .await
        .unwrap();

    let tx_status = wait_transaction(&http_url, tx_hash).await.unwrap().status();
    assert!(tx_status);
}

#[tokio::test]
async fn integration_bls_agg() -> Result<()> {
    // Start an Anvil instance
    let anvil = Anvil::new().spawn();
    let anvil_http_endpoint = anvil.endpoint().to_string();
    println!("anvil_http_endpoint: {}", anvil_http_endpoint);
    let anvil_ws_endpoint = anvil.ws_endpoint().to_string();
    println!("anvil_ws_endpoint: {}", anvil_ws_endpoint);
    // Retrieve deployed contract addresses (adjust with your util)
    let contract_addrs = get_contract_addresses(anvil_http_endpoint.clone()).await?;
    let signer_pk_str = ANVIL_FIRST_PRIVATE_KEY.to_string();
    let signer_pk = PrivateKeySigner::from_str(signer_pk_str.as_str()).expect("wrong key ");
    let wallet = EthereumWallet::from(signer_pk);

    // Build AVS registry clients
    let avs_registry_chain_reader = AvsRegistryChainReader::new(
        get_test_logger(),
        contract_addrs.registry_coordinator,
        contract_addrs.operator_state_retriever,
        anvil_http_endpoint.clone(),
    )
    .await?;
    let avs_registry_chain_writer = AvsRegistryChainWriter::build_avs_registry_chain_writer(
        get_test_logger(),
        anvil_http_endpoint.clone(),
        signer_pk_str,
        contract_addrs.registry_coordinator,
        contract_addrs.service_manager,
    )
    .await?;

    let bls_key_pair = BlsKeyPair::new(ANVIL_FIRST_PRIVATE_KEY.to_string())?;

    // Register operator on-chain
    let eth_client = Provider::<Http>::try_from(anvil_http_endpoint.clone())?;
    register_operator(
        &avs_registry_chain_writer,
        bls_key_pair.clone(),
        vec![0].into(),
        "localhost:8080".to_string(),
    )
    .await;

    let (operators_info_service, receiver) = OperatorInfoServiceInMemory::new(
        get_logger(),
        avs_registry_chain_reader.clone(),
        anvil_ws_endpoint,
    )
    .await?;

    let avs_registry_service = AvsRegistryServiceChainCaller::new(
        avs_registry_chain_reader.clone(),
        operators_info_service,
    );

    let quorum_numbers = vec![0];
    let socket = "localhost:8080".to_string();
    let salt_bytes = [0u8; 32];
    let salt: FixedBytes<32> = salt_bytes.into();

    // Compute expiry as 1 hour from now (Unix timestamp)
    let expiry_secs = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_secs()
        + 3600;
    let signature_expiry: U256 = expiry_secs
        .try_into()
        .expect("Failed to convert expiry to U256");

    avs_registry_chain_writer
        .register_operator_in_quorum_with_avs_registry_coordinator(
            bls_key_pair.clone(),
            salt,
            signature_expiry,
            quorum_numbers.clone().into(),
            socket,
        )
        .await?;

    // Start the node server to respond to certification requests
    let even_loving_certifier = EvenLovingCertifier {};
    let even_loving_node = BaseNode::new(
        Config {
            service_port: 8080,
            bls_key_pair: bls_key_pair.clone(),
        },
        even_loving_certifier,
    );

    spawn(async move {
        even_loving_node.start().await.unwrap();
    });

    let current_block_number = eth_client.get_block_number().await?;

    let operator_requester = GrpcOperatorRequester::new();

    let aggregator = AggregatorService::new(
        avs_registry_chain_reader.clone(),
        avs_registry_service.clone(),
        operator_requester.clone(),
    );

    let task_index = 0;
    let quorum_number = 0;
    let quorum_threshold_percentage = 100;
    let data = 69420i64.to_le_bytes().to_vec();
    let certificate_expiry = Duration::from_secs(10);

    let _ = aggregator
        .get_certificate(
            task_index,
            current_block_number
                .try_into()
                .expect("Failed to convert block number to u32"),
            quorum_number,
            quorum_threshold_percentage,
            data,
            certificate_expiry,
        )
        .await?;

    Ok(())
}
