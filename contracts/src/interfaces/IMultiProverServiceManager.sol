// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {IServiceManager} from "eigenlayer-middleware/interfaces/IServiceManager.sol";
import {BLSSignatureChecker} from "eigenlayer-middleware/BLSSignatureChecker.sol";

interface IMultiProverServiceManager is IServiceManager {
    /**
     * @notice Emitted when a state is confirmed
     */
    event StateConfirmed(
        uint256 indexed identifier,
        bytes metadata,
        bytes state
    );

    /**
     * @notice Emitted when the state confirmer is updated
     * @param previousConfirmer is the previous state confirmer
     * @param currentConfirmer is the current state confirmer
     */
    event StateConfirmerUpdated(address previousConfirmer, address currentConfirmer);

    /**
     * @notice The state proved by the prover
     * @param identifier identifier of the task that the prover is proving
     * @param metadata metadata of the task
     *                  - for example, if the task is to prove that a zk-rollup reached a certain state, the metadata could be the chainID and the target block number
     * @param state final state of the task
     * @param quorumNumbers each byte is a different quorum number
     * @param quorumThresholdPercentages each bytes is an amount less than 100 specifying the percentage of stake that has signed in the corresponding quorum in `quorumNumbers`
     * @param referenceBlockNumber the block number at which the stake information is being verified
     */
    struct StateHeader {
        uint256 identifier;
        bytes metadata;
        bytes state;
        bytes quorumNumbers;
        bytes quorumThresholdPercentages;
        uint32 referenceBlockNumber;
    }

    struct ReducedStateHeader {
        uint256 identifier;
        bytes metadata;
        bytes state;
        uint32 referenceBlockNumber;
    }

    function confirmState(
        StateHeader calldata stateHeader,
        BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external;
}