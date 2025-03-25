// SPDX-License-Identifier: UNLICENSED 
pragma solidity ^0.8.12;

import {BLSSignatureChecker} from "eigenlayer-middleware/src/BLSSignatureChecker.sol";
import {ISlashingRegistryCoordinator} from "eigenlayer-middleware/src/interfaces/ISlashingRegistryCoordinator.sol";

contract MinimalCertificateVerifier is BLSSignatureChecker {

    // CONSTANTS
    uint256 public constant DENOMINATOR = 1e18;
    uint256 public constant THRESHOLD = DENOMINATOR / 2;

    // STORAGE
    struct VerificationRecord {
        bytes inputData;
        bytes quorumNumbers;
        bytes32 responseHash;
        uint32 referenceBlockNumber;
        bytes32 signatoryRecordHash;
        QuorumStakeTotals quorumStakeTotals;
    }

    event CertificateVerified(bytes32 indexed responseHash, bytes inputData, uint256 recordIndex);
    
    VerificationRecord[] public verificationRecords;
    mapping(bytes32 => bool) public isCertificateVerified;

    constructor(
        ISlashingRegistryCoordinator __slashingRegistryCoordinator
    )
        BLSSignatureChecker(__slashingRegistryCoordinator)
    { }

    function verifyCertificate(
        bytes calldata inputData,
        bytes calldata response,
        bytes calldata quorumNumbers,
        uint32 referenceBlockNumber, 
        NonSignerStakesAndSignature calldata params
    ) external {
        bytes32 responseHash = keccak256(response);
        require(
            !isCertificateVerified[responseHash],
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

        emit CertificateVerified(responseHash, inputData, verificationRecords.length);

        verificationRecords.push(VerificationRecord(
            inputData,
            quorumNumbers,
            responseHash,
            referenceBlockNumber,
            signatoryRecordHash,
            quorumStakeTotals
        ));
        isCertificateVerified[responseHash] = true;
    }

    function totalCertificateCount() public view returns (uint256) {
        return verificationRecords.length;
    }
}
