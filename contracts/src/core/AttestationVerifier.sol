// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IAttestation} from "src/interfaces/IAttestation.sol";
import {IAttestationVerifier} from "src/interfaces/IAttestationVerifier.sol";

contract AttestationVerifier is IAttestationVerifier {

    IAttestation public dcapAttestation;

    constructor(address _attestationVerifierAddr) {
        dcapAttestation = IAttestation(_attestationVerifierAddr);
    }

    error INVALID_REPORT();
    error INVALID_REPORT_DATA();

    function verifyAttestation(bytes calldata _report) public returns (bytes memory) {
        (bool succ, bytes memory output) = dcapAttestation.verifyAndAttestOnChain(_report);
        if (!succ) {
            revert INVALID_REPORT();
        }

        if (output.length < 64) {
            revert INVALID_REPORT_DATA();
        }

        bytes memory reportData = new bytes(64);
        assembly {
            let start := add(add(output, 0x20), sub(mload(output), 128))
            mstore(add(reportData, 0x20), mload(start))
            mstore(add(reportData, 0x40), mload(add(start, 32)))
        }

        return reportData;
    }

    function verifyMrEnclave(bytes32) external view returns (bool) {
        return true;
    }

    function verifyMrSigner(bytes32) external view returns (bool) {
        return true;
    }

}
