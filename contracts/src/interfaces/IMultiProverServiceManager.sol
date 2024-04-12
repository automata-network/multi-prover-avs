//SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {IServiceManager} from "eigenlayer-middleware/interfaces/IServiceManager.sol";
import {BLSSignatureChecker} from "eigenlayer-middleware/BLSSignatureChecker.sol";

interface IMultiProverServiceManager is IServiceManager {
    /**
     * @notice Emitted when a state is confirmed
     */
    event StateConfirmed(
        uint256 indexed committeeId,
        bytes metadata,
        bytes state
    );

    /**
     * @notice Emitted when the state confirmer is updated
     * @param previousConfirmer is the previous state confirmer
     * @param currentConfirmer is the current state confirmer
     */
    event StateConfirmerUpdated(address previousConfirmer, address currentConfirmer);

    event PoAManagerUpdated(address previousPoaManager, address currentPoaManager);

    event CommitteeManagerUpdated(address previousCommitteeManager, address currentCommitteeManager);

    error NoPermission();

    error NotWhitelisted();

    error InvalidSender();

    error InvalidQuorumParam();

    error InvalidQuorum();

    error InsufficientThreshold();

    error ZeroAddr();

    error ZeroId();

    error CommitteeNotExist();

    error CommitteeExist();

    error TEEQuorumNotExist();

    error TEEQuorumExist();

    error TEEQuorumUsed();

    error QuorumNotInitialized();

    /**
     * @notice The state proved by the prover
     * @param committeeId the identifier of the committee
     * @param metadata metadata of the task
     *                  - for example, if the task is to prove that a zk-rollup reached a certain state, the metadata could be the chainID and the target block number
     * @param state final state of the task
     * @param quorumNumbers each byte is a different quorum number
     * @param quorumThresholdPercentages each bytes is an amount less than 100 specifying the percentage of stake that has signed in the corresponding quorum in `quorumNumbers`
     * @param referenceBlockNumber the block number at which the stake information is being verified
     */
    struct StateHeader {
        uint256 committeeId;
        bytes metadata;
        bytes state;
        bytes quorumNumbers;
        bytes quorumThresholdPercentages;
        uint32 referenceBlockNumber;
    }

    struct ReducedStateHeader {
        uint256 committeeId;
        bytes metadata;
        bytes state;
        uint32 referenceBlockNumber;
    }

    /**
     * @notice Enum of TEE types
     */
    enum TEE {
        NONE,
        INTEL_SGX,
        AWS_NITRO,
        AMD_SEV,
        ARM_TRUSTZONE
    }

    /**
     * @notice Struct representing a TEE quorum
     * @param teeType the type of TEE
     * @param quorumNumber the number of the corresponding quorum
     */
    struct TEEQuorum {
        TEE teeType;
        uint8 quorumNumber;
    }

    /**
     * @notice Struct representing a committee, which presents a set of operators that handle the same task
     * @param id the identifier of the committee
     * @param description the description of the committee
     * @param metadata the metadata of the committee, for example, the hash of binary used by the operators
     * @param teeQuorumNumbers the quorum numbers of the TEEs that are part of the committee, each byte is a different quorum number
     */
    struct Committee {
        uint256 id;
        string description;
        bytes metadata;
        bytes teeQuorumNumbers;
    }

    function confirmState(
        StateHeader calldata stateHeader,
        BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external;

    function addCommittee(Committee memory committee) external;

    function updateCommittee(Committee memory committee) external;

    function removeCommittee(uint256 committeeId) external;

    function addTEEQuorum(TEEQuorum memory teeQuorum) external;

    function removeTEEQuorum(uint8 quorumNumber) external;
}