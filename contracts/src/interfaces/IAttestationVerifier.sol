//SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

interface IAttestationVerifier {

    function verifyAttestation(bytes calldata _report) external payable returns (bytes memory);
    function verifyMrEnclave(bytes32 _mrEnclave) external view returns (bool);
    function verifyMrSigner(bytes32 _mrSigner) external view returns (bool);

    /**
     * @dev Estimates the fee for verifying the quote on-chain.
     * @param rawQuote The raw quote data.
     * @return The estimated fee.
     * @notice The actual fee is determined by multiplying the base fee with the gas price.
     */
    function estimateBaseFeeVerifyOnChain(bytes calldata rawQuote) external payable returns (uint256);

}
