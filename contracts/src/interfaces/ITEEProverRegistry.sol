//SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

/**
 * @notice interface for the TEE prover registry, which will
 *          - handle the registration/deregisteration of TEE provers, such as verifying the attestation of the TEE prover
 *          - handle the liveness challenge of TEE provers
 */
interface ITEEProverRegistry {

    /**
     * @notice Handles the registration of the TEE prover, will verify the attestation of the TEE prover
     * @param attestation the attestation of the TEE prover, which may varies depending on the platform of the TEE prover
     */
    function registerTEEProver(bytes memory attestation) external;

    /**
     * @notice Handles the deregistration of the TEE prover
     * @param signature the signature of the TEE prover to be deregistered
     */
    function deregisterTEEProver(bytes memory signature) external;

    /**
     * @notice Handles the liveness challenge of the TEE prover
     * @param challenge the challenge signed by the TEE prover to prove its liveness
     */
    function submitLivenessChallenge(bytes memory challenge) external;

    /**
     * @notice Checks the validity of the TEE prover, which includes the registration status and the liveness status
     * @param prover the address of the TEE prover
     */
    function checkTEEProverStatus(address prover) external view returns (bool valid);

}