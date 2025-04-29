//! End-to-end integration test for BLS aggregation

use alloy::network::{Ethereum, EthereumWallet, NetworkWallet};
use alloy::primitives::{Address, Bytes, FixedBytes, U256};
use alloy::signers::local::PrivateKeySigner;
use anyhow::Result;
use async_trait::async_trait;
use chrono::Local;
use eigensdk::client_avsregistry::reader::{AvsRegistryChainReader, AvsRegistryReader};
use eigensdk::client_avsregistry::writer::AvsRegistryChainWriter;
use eigensdk::client_elcontracts::reader::ELChainReader;
use eigensdk::crypto_bls::BlsKeyPair;
use eigensdk::logging::log_level::LogLevel;
use eigensdk::logging::{get_logger, get_test_logger, init_logger};
use eigensdk::services_avsregistry::chaincaller::AvsRegistryServiceChainCaller;
use eigensdk::services_operatorsinfo::operator_info::OperatorInfoService;
use eigensdk::services_operatorsinfo::operatorsinfo_inmemory::{
    OperatorInfoServiceError, OperatorInfoServiceInMemory,
};
use eigensdk::testing_utils::anvil_constants::{
    get_allocation_manager_address, get_avs_directory_address, get_delegation_manager_address,
    get_operator_state_retriever_address, get_registry_coordinator_address,
    get_rewards_coordinator_address, get_service_manager_address,
};
use eigensdk::testing_utils::transaction::wait_transaction;
use eigensdk::types::operator::{OperatorInfo, OperatorPubKeys};
use ethers::providers::{Http, Middleware, Provider};
use even_loving::EvenLovingCertifier;
use std::path::{Path, PathBuf};
use std::str::FromStr;
use std::time::{Duration, SystemTime, UNIX_EPOCH};
use teal::node::server::{BaseNode, Config};
use testcontainers::core::IntoContainerPort;
use testcontainers::core::WaitFor;
use testcontainers::runners::AsyncRunner;
use testcontainers::ImageExt;
use testcontainers::{ContainerAsync, GenericImage};
use tokio::spawn;
use tokio_util::sync::CancellationToken;

pub mod even_loving;

use teal::{aggregator::operator_requester::GrpcOperatorRequester, AggregatorService};

const ANVIL_FIRST_PRIVATE_KEY: &str =
    "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80";

struct ContractAddresses {
    service_manager: Address,
    allocation_manager: Address,
    delegation_manager: Address,
    rewards_coordinator: Address,
    avs_directory: Address,
    registry_coordinator: Address,
    operator_state_retriever: Address,
}

async fn get_contract_addresses(http_endpoint: String) -> Result<ContractAddresses> {
    let service_manager = get_service_manager_address(http_endpoint.clone()).await;
    let allocation_manager = get_allocation_manager_address(http_endpoint.clone()).await;
    let delegation_manager = get_delegation_manager_address(http_endpoint.clone()).await;
    let rewards_coordinator = get_rewards_coordinator_address(http_endpoint.clone()).await;
    let avs_directory = get_avs_directory_address(http_endpoint.clone()).await;
    let registry_coordinator = get_registry_coordinator_address(http_endpoint.clone()).await;
    let operator_state_retriever =
        get_operator_state_retriever_address(http_endpoint.clone()).await;

    Ok(ContractAddresses {
        service_manager,
        allocation_manager,
        delegation_manager,
        rewards_coordinator,
        avs_directory,
        registry_coordinator,
        operator_state_retriever,
    })
}

async fn register_operator(
    avs_writer: &AvsRegistryChainWriter,
    el_reader: &ELChainReader,
    service_manager: Address,
    operator_address: Address,
    operator_socket: String,
    bls_key_pair: BlsKeyPair,
    quorum_nums: Bytes,
    http_url: String,
) -> Result<()> {
    let salt_bytes = [0u8; 32];
    let salt: FixedBytes<32> = salt_bytes.into();

    let expiry_secs = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_secs()
        + 3600;
    let signature_expiry: U256 = expiry_secs
        .try_into()
        .expect("Failed to convert expiry to U256");

    let tx_hash = avs_writer
        .register_operator_in_quorum_with_avs_registry_coordinator(
            bls_key_pair,
            salt,
            signature_expiry,
            quorum_nums.clone(),
            operator_socket,
        )
        .await
        .unwrap();

    let tx = wait_transaction(&http_url, tx_hash).await.unwrap();
    assert!(tx.status(), "Transaction failed {:?}", tx);

    Ok(())
}

fn workspace_dir() -> PathBuf {
    let output = std::process::Command::new(env!("CARGO"))
        .arg("locate-project")
        .arg("--workspace")
        .arg("--message-format=plain")
        .output()
        .unwrap()
        .stdout;
    let cargo_path = Path::new(std::str::from_utf8(&output).unwrap().trim());
    cargo_path.parent().unwrap().to_path_buf()
}

const ANVIL_IMAGE: &str = "ghcr.io/foundry-rs/foundry";
const ANVIL_TAG: &str = "latest";

/// Copied from eigensdk-rs test utils so that we can pass our own state file
async fn start_anvil_with_state(
    state_path: &str,
) -> (ContainerAsync<GenericImage>, String, String) {
    let relative_path = PathBuf::from(state_path);
    let absolute_path = workspace_dir().join(relative_path);
    let absolute_path_str = absolute_path.to_str().unwrap();

    let container = GenericImage::new(ANVIL_IMAGE, ANVIL_TAG)
        .with_entrypoint("anvil")
        .with_wait_for(WaitFor::message_on_stdout("Listening on"))
        .with_exposed_port(8545u16.tcp())
        .with_mount(testcontainers::core::Mount::bind_mount(
            absolute_path_str,
            "/state.json",
        ))
        .with_cmd([
            "--host",
            "0.0.0.0",
            "--load-state",
            "/state.json",
            "--base-fee",
            "0",
            "--gas-price",
            "0",
            "--port",
            "8545",
        ])
        .start()
        .await
        .unwrap();

    let port = container
        .ports()
        .await
        .unwrap()
        .map_to_host_port_ipv4(8545u16.tcp())
        .unwrap();

    let http_endpoint = format!("http://localhost:{port}");
    let ws_endpoint = format!("ws://localhost:{port}");

    (container, http_endpoint, ws_endpoint)
}

#[derive(Clone)]
struct LocalOperatorInfoService {
    operator_socket: String,
    pub_keys: OperatorPubKeys,
}

#[async_trait]
impl OperatorInfoService for LocalOperatorInfoService {
    async fn get_operator_socket(
        &self,
        operator_address: Address,
    ) -> std::result::Result<Option<String>, OperatorInfoServiceError> {
        Ok(Some("localhost:8080".to_string()))
    }
    async fn get_operator_info(
        &self,
        operator_address: Address,
    ) -> std::result::Result<Option<OperatorPubKeys>, OperatorInfoServiceError> {
        Ok(Some(self.pub_keys.clone()))
    }
}

#[tokio::test]
async fn integration_bls_agg() -> Result<()> {
    init_logger(LogLevel::Debug);
    // Start an Anvil instance
    let (anvil, anvil_http_endpoint, anvil_ws_endpoint) =
        start_anvil_with_state("tests/contracts-deployed-anvil-state.json").await;
    println!("anvil_http_endpoint: {}", anvil_http_endpoint);
    println!("anvil_ws_endpoint: {}", anvil_ws_endpoint);
    // Retrieve deployed contract addresses (adjust with your util)
    let contract_addrs = get_contract_addresses(anvil_http_endpoint.clone()).await?;
    let signer_pk_str = ANVIL_FIRST_PRIVATE_KEY.to_string();
    let wallet = PrivateKeySigner::from_str(signer_pk_str.as_str()).expect("wrong key ");
    let operator_address = wallet.address();

    let el_chain_reader = ELChainReader::build(
        get_logger(),
        contract_addrs.delegation_manager,
        contract_addrs.avs_directory,
        contract_addrs.rewards_coordinator,
        &anvil_http_endpoint,
    )
    .await?;
    // Build AVS registry clients
    let avs_registry_chain_reader = AvsRegistryChainReader::new(
        get_logger(),
        contract_addrs.registry_coordinator,
        contract_addrs.operator_state_retriever,
        anvil_http_endpoint.clone(),
    )
    .await?;
    let avs_registry_chain_writer = AvsRegistryChainWriter::build_avs_registry_chain_writer(
        get_logger(),
        anvil_http_endpoint.clone(),
        signer_pk_str,
        contract_addrs.registry_coordinator,
        contract_addrs.service_manager,
    )
    .await?;

    let operator_socket = "localhost:8080".to_string();
    let bls_key_pair = BlsKeyPair::new(
        "1371012690269088913462269866874713266643928125698382731338806296762673180359922"
            .to_string(),
    )?;
    let eth_client = Provider::<Http>::try_from(anvil_http_endpoint.clone())?;

    register_operator(
        &avs_registry_chain_writer,
        &el_chain_reader,
        contract_addrs.service_manager,
        operator_address,
        operator_socket.clone(),
        bls_key_pair.clone(),
        vec![0].into(),
        anvil_http_endpoint,
    )
    .await?;

    let avs_registry_service = AvsRegistryServiceChainCaller::new(
        avs_registry_chain_reader.clone(),
        LocalOperatorInfoService {
            operator_socket: operator_socket.clone(),
            pub_keys: bls_key_pair.clone().into(),
        },
    );

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
        even_loving_node
            .start()
            .await
            .expect("Failed to start even loving node");
    });

    let operator_requester =
        GrpcOperatorRequester::new(format!("http://{}", operator_socket.clone()));

    let aggregator = AggregatorService::new(
        avs_registry_chain_reader.clone(),
        avs_registry_service.clone(),
        operator_requester.clone(),
    );

    let current_block_number = eth_client.get_block_number().await?;
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
