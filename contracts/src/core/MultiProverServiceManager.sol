// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {Pausable} from "eigenlayer-contracts/src/contracts/permissions/Pausable.sol";
import {IPauserRegistry} from "eigenlayer-contracts/src/contracts/interfaces/IPauserRegistry.sol";

import {ServiceManagerBase, IAVSDirectory} from "eigenlayer-middleware/ServiceManagerBase.sol";
import {BLSSignatureChecker} from "eigenlayer-middleware/BLSSignatureChecker.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/interfaces/IRegistryCoordinator.sol";
import {IStakeRegistry} from "eigenlayer-middleware/interfaces/IStakeRegistry.sol";

import {MultiProverServiceManagerStorage} from "./MultiProverServiceManagerStorage.sol";

contract MultiProverServiceManager is MultiProverServiceManagerStorage, ServiceManagerBase, BLSSignatureChecker, Pausable {

    modifier onlyStateConfirmer() {
        require(msg.sender == stateConfirmer, "MultiProverServiceManager.onlyStateConfirmer: caller is not the state confirmer");
        _;
    }

    constructor(
        IAVSDirectory __avsDirectory,
        IRegistryCoordinator __registryCoordinator,
        IStakeRegistry __stakeRegistry
    )
        BLSSignatureChecker(__registryCoordinator)
        ServiceManagerBase(__avsDirectory, __registryCoordinator, __stakeRegistry)
    {
        _disableInitializers();
    }

    function initialize(
        IPauserRegistry _pauserRegistry,
        uint256 _initialPausedStatus,
        address _initialOwner,
        address _stateConfirmer
    ) public initializer {
        _initializePauser(_pauserRegistry, _initialPausedStatus);
        _transferOwnership(_initialOwner);
        _setStateConfirmer(_stateConfirmer);
    }

    function confirmState(
        StateHeader calldata stateHeader,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external override onlyWhenNotPaused(PAUSED_SUBMIT_STATE) onlyStateConfirmer {
        // make sure the information needed to derive the non-signers and state transition is in calldata to avoid emitting events
        require(tx.origin == msg.sender, "MultiProverServiceManager.confirmState: state transition and nonsigner data must be in calldata");

        //make sure that the quorumNumbers and signedStakeForQuorums are of the same length
        require(
            stateHeader.quorumNumbers.length == stateHeader.quorumThresholdPercentages.length,
            "MultiProverServiceManager.confirmState: quorumNumbers and signedStakeForQuorums must be of the same length"
        );

        // calculate the hash of the state that operators are signing
        bytes32 reducedStateHeaderHash = _hashReducedStateHeader(_convertStateHeaderToReducedStateHeader(stateHeader));

        // check signatures
        // check the signature
        (
            QuorumStakeTotals memory quorumStakeTotals,
            bytes32 signatoryRecordHash
        ) = checkSignatures(
            reducedStateHeaderHash, 
            stateHeader.quorumNumbers, // use list of uint8s instead of uint256 bitmap to not iterate 256 times
            stateHeader.referenceBlockNumber, 
            nonSignerStakesAndSignature
        );

        // check that signatories own at least a threshold percentage of each quourm
        for (uint256 i = 0; i < stateHeader.quorumThresholdPercentages.length; i++) {
            // we don't check that the quorumThresholdPercentages are not >100 because a greater value would trivially fail the check, implying signed stake > total stake
            require(
                quorumStakeTotals.signedStakeForQuorum[i] * THRESHOLD_DENOMINATOR
                    >= quorumStakeTotals.totalStakeForQuorum[i] * uint8(stateHeader.quorumThresholdPercentages[i]),
                "MultiProverServiceManager.confirmState: signatories do not own at least threshold percentage of a quorum"
            );
        }

        uint32 taskIdMemory = taskId;
        bytes32 stateHeaderHash = _hashStateHeader(stateHeader);
        bytes32 taskMetadataHash = keccak256(abi.encodePacked(stateHeaderHash, signatoryRecordHash, uint32(block.number)));
        taskIdToMetadataHash[taskIdMemory] = taskMetadataHash;

        taskId = taskIdMemory + 1;

        emit StateConfirmed(stateHeader.identifier, stateHeader.metadata, stateHeader.state);
    }

    function setStateConfirmer(address _stateConfirmer) external onlyOwner {
        _setStateConfirmer(_stateConfirmer);
    }

    function _setStateConfirmer(address _stateConfirmer) internal {
        require(_stateConfirmer != address(0), "MultiProverServiceManager._setStateConfirmer: state confirmrt cannot be the zero address");
        address previousConfirmer = stateConfirmer;
        stateConfirmer = _stateConfirmer;
        emit StateConfirmerUpdated(previousConfirmer, _stateConfirmer);
    }

    function _hashStateHeader(StateHeader calldata stateHeader) internal pure returns (bytes32) {
        return keccak256(abi.encode(stateHeader));
    }

    // public ReducedStateHeader for go binding
    function _hashReducedStateHeader(ReducedStateHeader memory reducedStateHeader) public pure returns (bytes32) {
        return keccak256(abi.encode(reducedStateHeader));
    }

    function _convertStateHeaderToReducedStateHeader(StateHeader calldata stateHeader) internal pure 
        returns (ReducedStateHeader memory) 
    {
        return ReducedStateHeader({
            identifier: stateHeader.identifier,
            metadata: stateHeader.metadata,
            state: stateHeader.state,
            referenceBlockNumber: stateHeader.referenceBlockNumber
        });
    }
}