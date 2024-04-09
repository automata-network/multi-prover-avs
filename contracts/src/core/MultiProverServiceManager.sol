// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import {Pausable} from "eigenlayer-contracts/src/contracts/permissions/Pausable.sol";
import {IPauserRegistry} from "eigenlayer-contracts/src/contracts/interfaces/IPauserRegistry.sol";
import {ISignatureUtils} from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";

import {ServiceManagerBase, IAVSDirectory} from "eigenlayer-middleware/ServiceManagerBase.sol";
import {BLSSignatureChecker} from "eigenlayer-middleware/BLSSignatureChecker.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/interfaces/IRegistryCoordinator.sol";
import {IStakeRegistry} from "eigenlayer-middleware/interfaces/IStakeRegistry.sol";
import {IServiceManager} from "eigenlayer-middleware/interfaces/IServiceManager.sol";

import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {MultiProverServiceManagerStorage} from "./MultiProverServiceManagerStorage.sol";

contract MultiProverServiceManager is MultiProverServiceManagerStorage, ServiceManagerBase, BLSSignatureChecker, Pausable {
    using EnumerableSet for EnumerableSet.AddressSet;

    modifier onlyStateConfirmer() {
        require(msg.sender == stateConfirmer, "caller is not the state confirmer");
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
        address _stateConfirmer,
        bool _poaEnabled
    ) public initializer {
        _initializePauser(_pauserRegistry, _initialPausedStatus);
        _transferOwnership(_initialOwner);
        _setStateConfirmer(_stateConfirmer);
        poaEnabled = _poaEnabled;
    }

    function registerOperatorToAVS(
        address operator,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature) 
        public 
        override(ServiceManagerBase, IServiceManager)
        onlyWhenNotPaused(PAUSED_OPERTOR_REGISTRATION)
        onlyRegistryCoordinator 
    {
        require(!poaEnabled || operatorWhitelist.contains(operator), "operator not whitelisted");
        _avsDirectory.registerOperatorToAVS(operator, operatorSignature);
    }

    function deregisterOperatorFromAVS(address operator) 
        public 
        override(ServiceManagerBase, IServiceManager)
        onlyWhenNotPaused(PAUSED_OPERTOR_REGISTRATION)
        onlyRegistryCoordinator 
    {
        _avsDirectory.deregisterOperatorFromAVS(operator);
    }

    function confirmState(
        StateHeader calldata stateHeader,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external override onlyWhenNotPaused(PAUSED_SUBMIT_STATE) onlyStateConfirmer {
        // make sure the information needed to derive the non-signers and state transition is in calldata to avoid emitting events
        require(tx.origin == msg.sender, "state transition and nonsigner data must be in calldata");

        //make sure that the quorumNumbers and signedStakeForQuorums are of the same length
        require(
            stateHeader.quorumNumbers.length == stateHeader.quorumThresholdPercentages.length,
            "quorumNumbers and signedStakeForQuorums must be of the same length"
        );

        // make sure the committee exists
        require(committees[stateHeader.committeeId].id != 0, "committee does not exist");

        // make sure the quorums belong to the committee
        for (uint256 i = 0; i < stateHeader.quorumNumbers.length; i++) {
            require(
                quorumIdToCommitteeId[uint8(stateHeader.quorumNumbers[i])] == stateHeader.committeeId,
                "quorum does not belong to committee"
            );
        }

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
                "signatories do not own at least threshold percentage of a quorum"
            );
        }

        uint32 taskIdMemory = taskId;
        bytes32 stateHeaderHash = _hashStateHeader(stateHeader);
        bytes32 taskMetadataHash = keccak256(abi.encodePacked(stateHeaderHash, signatoryRecordHash, uint32(block.number)));
        taskIdToMetadataHash[taskIdMemory] = taskMetadataHash;

        taskId = taskIdMemory + 1;

        emit StateConfirmed(stateHeader.committeeId, stateHeader.metadata, stateHeader.state);
    }

    function addCommittee(Committee memory committee) external onlyOwner {
        require(committee.id != 0, "committee id cannot be 0");
        require(committees[committee.id].id == 0, "committee already exists");
        for (uint256 i = 0; i < committee.teeQuorumNumbers.length; i++) {
            uint8 teeQuorumNumber = uint8(committee.teeQuorumNumbers[i]);
            require(teeQuorums[teeQuorumNumber].teeType != TEE.NONE, "tee quorum does not exist");
            require(quorumIdToCommitteeId[teeQuorumNumber] == 0, "tee quorum is used by another committee");
            quorumIdToCommitteeId[teeQuorumNumber] = committee.id;
        }

        committees[committee.id] = committee;
    }

    function updateCommittee(Committee memory committee) external onlyOwner {
        require(committees[committee.id].id != 0, "MultiProverServiceManager.updateCommittee: committee does not exist");
        for (uint256 i = 0; i < committee.teeQuorumNumbers.length; i++) {
            uint8 teeQuorumNumber = uint8(committee.teeQuorumNumbers[i]);
            require(teeQuorums[teeQuorumNumber].teeType != TEE.NONE, "tee quorum does not exist");
            require(quorumIdToCommitteeId[teeQuorumNumber] == 0 || quorumIdToCommitteeId[teeQuorumNumber] == committee.id, "MultiProverServiceManager.updateCommittee: tee quorum is used by another committee");
            quorumIdToCommitteeId[teeQuorumNumber] = committee.id;
        }

        committees[committee.id] = committee;
    }

    function removeCommittee(uint256 committeeId) external onlyOwner {
        require(committees[committeeId].id != 0, "committee does not exist");
        bytes memory teeQuorumNumbers = committees[committeeId].teeQuorumNumbers;
        for (uint256 i = 0; i < teeQuorumNumbers.length; i++) {
            quorumIdToCommitteeId[uint8(teeQuorumNumbers[i])] = 0;
        }

        delete committees[committeeId];
    }

    function addTEEQuorum(TEEQuorum memory teeQuorum) external onlyOwner {
        require(teeQuorums[teeQuorum.quorumNumber].teeType == TEE.NONE, "tee quorum already exists");
        require(_stakeRegistry.getTotalStakeHistoryLength(teeQuorum.quorumNumber) != 0, "MultiProverServiceManager.addTEEQuorum: quorum not initialized");

        teeQuorums[teeQuorum.quorumNumber] = teeQuorum;
    }

    function removeTEEQuorum(uint8 quorumNumber) external onlyOwner {
        require(teeQuorums[quorumNumber].teeType != TEE.NONE, "tee quorum does not exist");
        require(quorumIdToCommitteeId[quorumNumber] == 0, "tee quorum is in use");

        delete teeQuorums[quorumNumber];
    }

    function setStateConfirmer(address _stateConfirmer) external onlyOwner {
        _setStateConfirmer(_stateConfirmer);
    }

    function enablePoA() external onlyOwner {
        require(!poaEnabled, "PoA already enabled");
        poaEnabled = true;
    }

    function disablePoA() external onlyOwner {
        require(poaEnabled, "PoA already disabled");
        poaEnabled = false;
    }

    function isPoAEnabled() external view returns (bool) {
        return poaEnabled;
    }

    function whitelistOperator(address operator) external onlyOwner {
        require(operator != address(0), "zero address");
        require(!operatorWhitelist.contains(operator), "operator already whitelisted");
        operatorWhitelist.add(operator);
    }

    function blacklistOperator(address operator) external onlyOwner {
        require(operator != address(0), "zero address");
        require(operatorWhitelist.contains(operator), "operator not whitelisted");
        operatorWhitelist.remove(operator);
    }

    function isOperatorWhitelisted(address operator) external view returns (bool) {
        return operatorWhitelist.contains(operator);
    }

    function _setStateConfirmer(address _stateConfirmer) internal {
        require(_stateConfirmer != address(0), "zero address");
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
            committeeId: stateHeader.committeeId,
            metadata: stateHeader.metadata,
            state: stateHeader.state,
            referenceBlockNumber: stateHeader.referenceBlockNumber
        });
    }
}
