// SPDX-License-Identifier: UNLICENSED 
pragma solidity ^0.8.12;

import {BLSSignatureChecker} from "eigenlayer-middleware/BLSSignatureChecker.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/interfaces/IRegistryCoordinator.sol";

contract MinimalCertificateVerifier is BLSSignatureChecker {

    // CONSTANTS
    uint256 public constant DENOMINATOR = 1e18;
    uint256 public constant THRESHOLD = DENOMINATOR / 2;

    // STORAGE
    struct VerificationRecord { 
        bytes quorumNumbers;
        uint32 referenceBlockNumber;
        bytes32 signatoryRecordHash;
        QuorumStakeTotals quorumStakeTotals;
    }
    
    mapping(bytes32 => VerificationRecord) public verificationRecords;

    constructor(
        IRegistryCoordinator __registryCoordinator
    )
        BLSSignatureChecker(__registryCoordinator)
    { }

    function verifyCertificate(
        bytes calldata response,
        bytes calldata quorumNumbers,
        uint32 referenceBlockNumber, 
        NonSignerStakesAndSignature calldata params
    ) external {
        bytes32 responseHash = keccak256(response);
        require(
            verificationRecords[responseHash].referenceBlockNumber == 0,
            "Certificate already verified"
        );

        (
            QuorumStakeTotals memory quorumStakeTotals,
            bytes32 signatoryRecordHash
        ) = checkSignatures(
            responseHash, 
            quorumNumbers, // use list of uint8s instead of uint256 bitmap to not iterate 256 times
            referenceBlockNumber, 
            params
        );

        for (uint256 i = 0; i < quorumStakeTotals.signedStakeForQuorum.length; i++) {
            require(
                quorumStakeTotals.signedStakeForQuorum[i] * DENOMINATOR >
                quorumStakeTotals.totalStakeForQuorum[i] * THRESHOLD,
                "Threshold not met"
            );
        }

        verificationRecords[responseHash] = VerificationRecord(
            quorumNumbers,
            referenceBlockNumber,
            signatoryRecordHash,
            quorumStakeTotals
        );
    }
}
