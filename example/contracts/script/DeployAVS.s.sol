// SPDX-License-Identifier: UNLICENSED 
pragma solidity ^0.8.12;

import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ITransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {PauserRegistry} from "eigenlayer-contracts/src/contracts/permissions/PauserRegistry.sol";
import {EmptyContract} from "eigenlayer-contracts/src/test/mocks/EmptyContract.sol";
import {IDelegationManager} from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import {IAVSDirectory} from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import {IRewardsCoordinator} from "eigenlayer-contracts/src/contracts/interfaces/IRewardsCoordinator.sol";
import {IAllocationManager} from "eigenlayer-contracts/src/contracts/interfaces/IAllocationManager.sol";
import {IAVSRegistrar} from "eigenlayer-contracts/src/contracts/interfaces/IAVSRegistrar.sol";
import {IPermissionController} from "eigenlayer-contracts/src/contracts/interfaces/IPermissionController.sol";


import {BLSApkRegistry} from "eigenlayer-middleware/src/BLSApkRegistry.sol";
import {SlashingRegistryCoordinator} from "eigenlayer-middleware/src/SlashingRegistryCoordinator.sol";
import {OperatorStateRetriever} from "eigenlayer-middleware/src/OperatorStateRetriever.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {IndexRegistry} from "eigenlayer-middleware/src/IndexRegistry.sol";
import {IIndexRegistry} from "eigenlayer-middleware/src/interfaces/IIndexRegistry.sol";
import {StakeRegistry, IStrategy} from "eigenlayer-middleware/src/StakeRegistry.sol";
import {IStakeRegistry, IStakeRegistryTypes} from "eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import {IServiceManager} from "eigenlayer-middleware/src/interfaces/IServiceManager.sol";
import {IBLSApkRegistry} from "eigenlayer-middleware/src/interfaces/IBLSApkRegistry.sol";
import {ServiceManagerBase} from "eigenlayer-middleware/src/ServiceManagerBase.sol";
import {ISocketRegistry, SocketRegistry} from "eigenlayer-middleware/src/SocketRegistry.sol";
import {IPauserRegistry} from "eigenlayer-contracts/src/contracts/interfaces/IPauserRegistry.sol";
import {ISlashingRegistryCoordinator, ISlashingRegistryCoordinatorTypes} from "eigenlayer-middleware/src/interfaces/ISlashingRegistryCoordinator.sol";

import {MinimalServiceManager} from "../src/MinimalServiceManager.sol";
import {MinimalCertificateVerifier} from "../src/MinimalCertificateVerifier.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";


contract DeployAVS is Script, Test {
    // Core contracts
    ProxyAdmin public avsProxyAdmin;
    PauserRegistry public avsPauserReg;
    EmptyContract public emptyContract;


    // Middleware contracts
    BLSApkRegistry public apkRegistry;
    IServiceManager public serviceManager;
    MinimalCertificateVerifier public certificateVerifier;
    SlashingRegistryCoordinator public slashingRegistryCoordinator;
    IIndexRegistry public indexRegistry;
    IStakeRegistry public stakeRegistry;
    ISocketRegistry public socketRegistry;
    OperatorStateRetriever public operatorStateRetriever;

    // Implementation contracts
    BLSApkRegistry public apkRegistryImplementation;
    IServiceManager public serviceManagerImplementation;
    MinimalCertificateVerifier public certificateVerifierImplementation;
    ISlashingRegistryCoordinator public slashingRegistryCoordinatorImplementation;
    IIndexRegistry public indexRegistryImplementation;
    IStakeRegistry public stakeRegistryImplementation;
    ISocketRegistry public socketRegistryImplementation;

    struct EigenlayerDeployment {
        address allocationManager;
        address delegationManager;
        address permissionController;
        address rewardsCoordinator;
        address avsDirectory;
    }
    
    function run(
        string memory inputConfigPath,
        uint256 maxOperatorCount,
        IStrategy[] memory strategies
    ) public virtual {
        // read the json file
        string memory inputConfig = vm.readFile(inputConfigPath);
        EigenlayerDeployment memory eigenlayerDeployment = EigenlayerDeployment({
            allocationManager: stdJson.readAddress(inputConfig, ".allocationManager"),
            delegationManager: stdJson.readAddress(inputConfig, ".delegationManager"),
            permissionController: stdJson.readAddress(inputConfig, ".permissionController"),
            rewardsCoordinator: stdJson.readAddress(inputConfig, ".rewardsCoordinator"),
            avsDirectory: stdJson.readAddress(inputConfig, ".avsDirectory")
        });


        emit log_named_address("allocation manager", eigenlayerDeployment.allocationManager);
        emit log_named_address("delegation manager", eigenlayerDeployment.delegationManager);
        emit log_named_address("permission controller", eigenlayerDeployment.permissionController);
        emit log_named_address("rewards coordinator", eigenlayerDeployment.rewardsCoordinator);
        emit log_named_address("avs directory", eigenlayerDeployment.avsDirectory);

        // only a lower bound for the deployment block number
        uint256 deploymentBlock = block.number;
        vm.startBroadcast();
        // deploy proxy admin for ability to upgrade proxy contracts
        avsProxyAdmin = new ProxyAdmin();

        // deploy pauser registry
        {
            address[] memory pausers = new address[](1);
            pausers[0] = msg.sender;
            avsPauserReg = new PauserRegistry(pausers, msg.sender);
        }

        emptyContract = new EmptyContract();
        
        // Deploy upgradeable proxy contracts pointing to empty contract initially
        serviceManager = ServiceManagerBase(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(avsProxyAdmin), ""))
        );

        certificateVerifier = MinimalCertificateVerifier(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(avsProxyAdmin), ""))
        );

        slashingRegistryCoordinator = SlashingRegistryCoordinator(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(avsProxyAdmin), ""))
        );

        indexRegistry = IIndexRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(avsProxyAdmin), ""))
        );

        stakeRegistry = IStakeRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(avsProxyAdmin), ""))
        );

        apkRegistry = BLSApkRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(avsProxyAdmin), ""))
        );

        socketRegistry = ISocketRegistry(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(avsProxyAdmin), ""))
        );

        // Deploy implementations and upgrade proxies
        indexRegistryImplementation = new IndexRegistry(
            slashingRegistryCoordinator
        );

        avsProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(indexRegistry))),
            address(indexRegistryImplementation)
        );

        stakeRegistryImplementation = new StakeRegistry(
            slashingRegistryCoordinator,
            IDelegationManager(eigenlayerDeployment.delegationManager),
            IAVSDirectory(eigenlayerDeployment.avsDirectory),
            IAllocationManager(eigenlayerDeployment.allocationManager)
        );

        avsProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(stakeRegistry))),
            address(stakeRegistryImplementation)
        );

        apkRegistryImplementation = new BLSApkRegistry(
            slashingRegistryCoordinator
        );

        avsProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(apkRegistry))),
            address(apkRegistryImplementation)
        );

        socketRegistryImplementation = new SocketRegistry(slashingRegistryCoordinator);

        avsProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(socketRegistry))),
            address(socketRegistryImplementation)
        );

        serviceManagerImplementation = new MinimalServiceManager(
            IAVSDirectory(eigenlayerDeployment.avsDirectory),
            IRewardsCoordinator(eigenlayerDeployment.rewardsCoordinator),
            slashingRegistryCoordinator,
            stakeRegistry,
            IPermissionController(address(eigenlayerDeployment.permissionController)),
            IAllocationManager(eigenlayerDeployment.allocationManager)
        );

        // Initialize ServiceManagerBase
        avsProxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(payable(address(serviceManager))),
            address(serviceManagerImplementation),
            abi.encodeWithSelector(
                MinimalServiceManager.initialize.selector,
                msg.sender,
                msg.sender
            )
        );

        slashingRegistryCoordinatorImplementation = new SlashingRegistryCoordinator(
                stakeRegistry,
                apkRegistry,
                indexRegistry,
                socketRegistry,
                IAllocationManager(eigenlayerDeployment.allocationManager),
                avsPauserReg
            );

        {
            ISlashingRegistryCoordinatorTypes.OperatorSetParam[] memory operatorSetParams = new ISlashingRegistryCoordinatorTypes.OperatorSetParam[](strategies.length);
            for (uint i = 0; i < strategies.length; i++) {
                operatorSetParams[i] = ISlashingRegistryCoordinatorTypes.OperatorSetParam({
                    maxOperatorCount: uint32(maxOperatorCount),
                    kickBIPsOfOperatorStake: 11000,
                    kickBIPsOfTotalStake: 1001
                });
            }

            uint96[] memory minimumStakeForQuourm = new uint96[](strategies.length);
            for (uint i = 0; i < strategies.length; i++) {
                minimumStakeForQuourm[i] = 1;
            }
            IStakeRegistryTypes.StrategyParams[][] memory strategyAndWeightingMultipliers = new IStakeRegistryTypes.StrategyParams[][](strategies.length);
            for (uint i = 0; i < strategies.length; i++) {
                strategyAndWeightingMultipliers[i] = new IStakeRegistryTypes.StrategyParams[](1);
                strategyAndWeightingMultipliers[i][0] = IStakeRegistryTypes.StrategyParams({
                    strategy: strategies[i],
                    multiplier: 1 ether
                });
            }

            avsProxyAdmin.upgradeAndCall(
                ITransparentUpgradeableProxy(payable(address(slashingRegistryCoordinator))),
                address(slashingRegistryCoordinatorImplementation),
                abi.encodeWithSelector(
                    SlashingRegistryCoordinator.initialize.selector,
                    msg.sender, // initial owner
                    msg.sender, // churn approver
                    msg.sender, // ejector
                    0, // initial paused status
                    address(serviceManager) // accountIdentifier
                )
            );

            // set AVS Registrar on AllocationManager to SlashingRegistryCoordinator
            serviceManager.setAppointee(
                msg.sender,
                eigenlayerDeployment.allocationManager,
                IAllocationManager(eigenlayerDeployment.allocationManager).setAVSRegistrar.selector
            );

            IAllocationManager(eigenlayerDeployment.allocationManager).setAVSRegistrar(
                address(serviceManager),
                IAVSRegistrar(slashingRegistryCoordinator)
            );

             // give slashingregistrycoordindator permission to createTotalDelegatedStakeQuorum
             serviceManager.setAppointee(
                address(slashingRegistryCoordinator),
                eigenlayerDeployment.allocationManager,
                IAllocationManager(eigenlayerDeployment.allocationManager).createOperatorSets.selector
             );

            for (uint i = 0; i < strategies.length; i++) {
                slashingRegistryCoordinator.createSlashableStakeQuorum(
                    operatorSetParams[i],
                    minimumStakeForQuourm[i],
                    strategyAndWeightingMultipliers[i],
                    1
                );
            }   
        }

        certificateVerifierImplementation = new MinimalCertificateVerifier(
            slashingRegistryCoordinator
        );

        avsProxyAdmin.upgrade(
            ITransparentUpgradeableProxy(payable(address(certificateVerifier))),
            address(certificateVerifierImplementation)
        );

        operatorStateRetriever = new OperatorStateRetriever();

        vm.stopBroadcast();

        string memory output = "deployment";
        vm.serializeAddress(output, "serviceManager", address(serviceManager));
        vm.serializeAddress(output, "certificateVerifier", address(certificateVerifier));
        vm.serializeAddress(output, "slashingRegistryCoordinator", address(slashingRegistryCoordinator));
        vm.serializeAddress(output, "indexRegistry", address(indexRegistry));
        vm.serializeAddress(output, "stakeRegistry", address(stakeRegistry));
        vm.serializeAddress(output, "apkRegistry", address(apkRegistry));
        vm.serializeAddress(output, "socketRegistry", address(socketRegistry));
        vm.serializeAddress(output, "operatorStateRetriever", address(operatorStateRetriever));
        vm.serializeAddress(output, "avsProxyAdmin", address(avsProxyAdmin));
        vm.serializeAddress(output, "avsPauserReg", address(avsPauserReg));
        address[] memory strategyAddresses = new address[](strategies.length);
        for (uint i = 0; i < strategies.length; i++) {
            strategyAddresses[i] = address(strategies[i]);
        }
        vm.serializeAddress(output, "strategies", strategyAddresses);
        vm.serializeUint(output, "deploymentBlock", deploymentBlock);

        string memory finalJson = vm.serializeString(output, "object", output);

        vm.createDir("./script/output", true);
        vm.writeJson(finalJson, "./script/output/avs_deploy_output.json");     
    }
}
