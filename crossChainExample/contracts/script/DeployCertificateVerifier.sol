// SPDX-License-Identifier: UNLICENSED 
pragma solidity ^0.8.12;

import {ProxyAdmin} from "eigenlayer-middleware/lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "eigenlayer-middleware/lib/openzeppelin-contracts/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ITransparentUpgradeableProxy} from "eigenlayer-middleware/lib/openzeppelin-contracts/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {EmptyContract} from "eigenlayer-core/test/mocks/EmptyContract.sol";

import {BLSCertificateVerifier} from "eigenlayer-middleware/src/crossChain/BLSCertificateVerifier.sol";

import {OperatorSet} from "eigenlayer-core/contracts/libraries/OperatorSetLib.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";

contract DeployCertificateVerifier is Script, Test {
    ProxyAdmin public proxyAdmin;
    EmptyContract public emptyContract;

    BLSCertificateVerifier public certificateVerifier;
    BLSCertificateVerifier public certificateVerifierImplementation;

    struct AVSDeployment {
        address serviceManager;
    }
    

    function run(
        string memory avsConfigPath,
        uint32 operatorSetId,
        uint32 maxTableStaleness
    ) public {
        // read the json file
        string memory avsConfig = vm.readFile(avsConfigPath);
        AVSDeployment memory avsDeployment = AVSDeployment({
            serviceManager: stdJson.readAddress(avsConfig, ".serviceManager")
        });

        uint256 deploymentBlock = block.number;
        vm.startBroadcast();

        OperatorSet memory operatorSet = OperatorSet({
            avs: avsDeployment.serviceManager,
            id: operatorSetId
        });

        // Deploy BLSCertificateVerifier
        certificateVerifier = new BLSCertificateVerifier(
            operatorSet,
            msg.sender,
            maxTableStaleness
        );

        // TODO: make this upgradeable
        // // Deploy ProxyAdmin
        // proxyAdmin = new ProxyAdmin();

        // emptyContract = new EmptyContract();

        // // Deploy proxy
        // certificateVerifier = new BLSCertificateVerifier(
        //     address (new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        // );

        // // Deploy implementation
        // certificateVerifierImplementation = new BLSCertificateVerifier(
        //     address (new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        // );

        // Upgrade Proxy
        // proxyAdmin.upgradeAndCall(
        //     ITransparentUpgradeableProxy(payable(address(certificateVerifier))),
        //     address(certificateVerifierImplementation),
        //     abi.encodeWithSelector(
        //         BLSCertificateVerifier.initialize.selector,
        //         operatorSet,
        //         msg.sender,
        //         maxTableStaleness,
        //         msg.sender
        //     )
        // );

        vm.stopBroadcast();

        // TODO: get the formatting right
        string memory output = "crossChainDeployment";
        vm.serializeAddress(output, "certificateVerifier", address(certificateVerifier));
        vm.serializeUint(output, "deploymentBlock", deploymentBlock);
        vm.serializeUint(output, "chainID", block.chainid);

        string memory finalJson = vm.serializeString(output, "object", output);

        vm.createDir("./script/output", true);
        vm.writeJson(finalJson, "./script/output/certificateVerifier_deploy_output.json");
    }

}
