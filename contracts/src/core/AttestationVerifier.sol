// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IAttestation} from "../interfaces/IAttestation.sol";
import {IAttestationVerifier} from "../interfaces/IAttestationVerifier.sol";

contract AttestationVerifier is IAttestationVerifier {

    IAttestation public dcapAttestation;

    constructor(address _attestationVerifierAddr) {
        dcapAttestation = IAttestation(_attestationVerifierAddr);
    }

    error INVALID_REPORT();
    error INVALID_REPORT_DATA();

    function verifyAttestation(bytes calldata _report) public payable returns (bytes memory) {
        (bool succ, bytes memory output) = dcapAttestation.verifyAndAttestOnChain{value: msg.value}(_report);
        if (!succ) {
            revert INVALID_REPORT();
        }

        if (output.length < 64) {
            revert INVALID_REPORT_DATA();
        }

        // tee = output[2:6]
        bytes4 tee;
        assembly {
            let start := add(add(output, 0x20), 2)
            tee := mload(start)
        }

        bytes memory reportData = new bytes(64);
        if (tee == 0x00000000) {
            // sgx, reportData = output[333:397]
            assembly {
                let start := add(add(output, 0x20), 333) // 13 + 384 - 64
                mstore(add(reportData, 0x20), mload(start))
                mstore(add(reportData, 0x40), mload(add(start, 32)))
            }
        } else {
            // tdx, reportData = output[533:597]
            assembly {
                let start := add(add(output, 0x20), 533) // 13 + 584 - 64
                mstore(add(reportData, 0x20), mload(start))
                mstore(add(reportData, 0x40), mload(add(start, 32)))
            }
        }
        return reportData;
    }

    function verifyMrEnclave(bytes32) external view returns (bool) {
        return true;
    }

    function verifyMrSigner(bytes32) external view returns (bool) {
        return true;
    }

    /**
     * @dev Estimates the fee for verifying the quote on-chain.
     * @param rawQuote The raw quote data.
     * @return The estimated fee.
     * @notice The actual fee is determined by multiplying the base fee with the gas price.
     */
    function estimateBaseFeeVerifyOnChain(bytes calldata rawQuote) external payable returns (uint256) {
        uint16 bp = dcapAttestation.getBp();
        uint256 gasBefore = gasleft();
        dcapAttestation.verifyAndAttestOnChain{value: msg.value}(rawQuote);
        uint256 gasAfter = gasleft();
        return (gasBefore - gasAfter) * bp / 10000;
    }
}
