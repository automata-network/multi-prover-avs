//SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {Pausable} from "eigenlayer-contracts/src/contracts/permissions/Pausable.sol";
import {IPauserRegistry} from "eigenlayer-contracts/src/contracts/interfaces/IPauserRegistry.sol";
import {ISignatureUtils} from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";

import {ServiceManagerBase, IAVSDirectory} from "eigenlayer-middleware/ServiceManagerBase.sol";
import {BLSSignatureChecker} from "eigenlayer-middleware/BLSSignatureChecker.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/interfaces/IRegistryCoordinator.sol";
import {IStakeRegistry} from "eigenlayer-middleware/interfaces/IStakeRegistry.sol";
import {IServiceManager} from "eigenlayer-middleware/interfaces/IServiceManager.sol";

import {MultiProverServiceManagerStorage} from "./MultiProverServiceManagerStorage.sol";

contract MultiProverServiceManager is MultiProverServiceManagerStorage, ServiceManagerBase, BLSSignatureChecker, Pausable {

    modifier onlyStateConfirmer() {
        if (msg.sender != stateConfirmer) {
            revert NoPermission();
        }
        _;
    }

    modifier onlyPoAManager() {
        if (msg.sender != poaManager) {
            revert NoPermission();
        }
        _;
    }

    modifier onlyCommitteeManager() {
        if (msg.sender != committeeManager) {
            revert NoPermission();
        }
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
        address _poaManager,
        address _committeeManager,
        bool _poaEnabled
    ) public initializer {
        _initializePauser(_pauserRegistry, _initialPausedStatus);
        __ServiceManagerBase_init(_initialOwner);
        _setStateConfirmer(_stateConfirmer);
        _setPoAManager(_poaManager);
        _setCommitteeManager(_committeeManager);
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
        if (poaEnabled && !operatorWhitelist[operator]) {
            revert NotWhitelisted();
        }
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
        if (tx.origin != msg.sender) {
            revert InvalidSender();
        }

        //make sure that the quorumNumbers and signedStakeForQuorums are of the same length
        if (stateHeader.quorumNumbers.length != stateHeader.quorumThresholdPercentages.length) {
            revert InvalidQuorumParam();
        }

        // make sure the committee exists
        if (committees[stateHeader.committeeId].id == 0) {
            revert CommitteeNotExist();
        }

        // make sure the quorums belong to the committee
        for (uint256 i = 0; i < stateHeader.quorumNumbers.length; i++) {
            if (quorumIdToCommitteeId[uint8(stateHeader.quorumNumbers[i])] != stateHeader.committeeId) {
                revert InvalidQuorum();
            }
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
            if (quorumStakeTotals.signedStakeForQuorum[i] * THRESHOLD_DENOMINATOR 
                    < quorumStakeTotals.totalStakeForQuorum[i] * uint8(stateHeader.quorumThresholdPercentages[i])) {
                revert InsufficientThreshold();
            }
        }

        uint32 taskIdMemory = taskId;
        bytes32 stateHeaderHash = _hashStateHeader(stateHeader);
        bytes32 taskMetadataHash = keccak256(abi.encodePacked(stateHeaderHash, signatoryRecordHash, uint32(block.number)));
        taskIdToMetadataHash[taskIdMemory] = taskMetadataHash;

        taskId = taskIdMemory + 1;

        emit StateConfirmed(stateHeader.committeeId, stateHeader.metadata, stateHeader.state);
    }

    function addCommittee(Committee memory committee) external onlyCommitteeManager {
        if (committee.id == 0) {
            revert ZeroId();
        }
        if (committees[committee.id].id != 0) {
            revert CommitteeExist();
        }
        for (uint256 i = 0; i < committee.teeQuorumNumbers.length; i++) {
            uint8 teeQuorumNumber = uint8(committee.teeQuorumNumbers[i]);
            if (teeQuorums[teeQuorumNumber].teeType == TEE.NONE) {
                revert TEEQuorumNotExist();
            }
            if (quorumIdToCommitteeId[teeQuorumNumber] != 0) {
                revert TEEQuorumUsed();
            }
            quorumIdToCommitteeId[teeQuorumNumber] = committee.id;
        }

        committees[committee.id] = committee;
    }

    function updateCommittee(Committee memory committee) external onlyCommitteeManager {
        if (committees[committee.id].id == 0) {
            revert CommitteeNotExist();
        }

        _removeCommittee(committee.id);
        for (uint256 i = 0; i < committee.teeQuorumNumbers.length; i++) {
            uint8 teeQuorumNumber = uint8(committee.teeQuorumNumbers[i]);
            if (teeQuorums[teeQuorumNumber].teeType == TEE.NONE) {
                revert TEEQuorumNotExist();
            }
            if (quorumIdToCommitteeId[teeQuorumNumber] != 0) {
                revert TEEQuorumUsed();
            }
            quorumIdToCommitteeId[teeQuorumNumber] = committee.id;
        }

        committees[committee.id] = committee;
    }

    function removeCommittee(uint256 committeeId) external onlyCommitteeManager {
        _removeCommittee(committeeId);
    }

    function _removeCommittee(uint256 committeeId) internal {
        if (committees[committeeId].id == 0) {
            revert CommitteeNotExist();
        }
        bytes memory teeQuorumNumbers = committees[committeeId].teeQuorumNumbers;
        for (uint256 i = 0; i < teeQuorumNumbers.length; i++) {
            quorumIdToCommitteeId[uint8(teeQuorumNumbers[i])] = 0;
        }

        delete committees[committeeId];
    }

    function addTEEQuorum(TEEQuorum memory teeQuorum) external onlyCommitteeManager {
        if (teeQuorums[teeQuorum.quorumNumber].teeType != TEE.NONE) {
            revert TEEQuorumExist();
        }
        if (teeQuorum.teeType == TEE.NONE) {
            revert InvalidQuorumParam();
        }
        if (_stakeRegistry.getTotalStakeHistoryLength(teeQuorum.quorumNumber) == 0) {
            revert QuorumNotInitialized();
        }

        teeQuorums[teeQuorum.quorumNumber] = teeQuorum;
    }

    function removeTEEQuorum(uint8 quorumNumber) external onlyCommitteeManager {
        if (teeQuorums[quorumNumber].teeType == TEE.NONE) {
            revert TEEQuorumNotExist();
        }
        if (quorumIdToCommitteeId[quorumNumber] != 0) {
            revert TEEQuorumUsed();
        }
        delete teeQuorums[quorumNumber];
    }

    function setStateConfirmer(address _stateConfirmer) external onlyOwner {
        _setStateConfirmer(_stateConfirmer);
    }

    function setPoAManager(address _poaManager) external onlyOwner {
        _setPoAManager(_poaManager);
    }

    function setCommitteeManager(address _committeeManager) external onlyOwner {
        _setCommitteeManager(_committeeManager);
    }

    function enablePoA() external onlyPoAManager {
        poaEnabled = true;
    }

    function disablePoA() external onlyPoAManager {
        poaEnabled = false;
    }

    function isPoAEnabled() external view returns (bool) {
        return poaEnabled;
    }

    function whitelistOperator(address operator) external onlyPoAManager {
        if (operator == address(0)) {
            revert ZeroAddr();
        }
        operatorWhitelist[operator] = true;
    }

    function blacklistOperator(address operator) external onlyPoAManager {
        if (operator == address(0)) {
            revert ZeroAddr();
        }
        operatorWhitelist[operator] = false;
    }

    function isOperatorWhitelisted(address operator) external view returns (bool) {
        return operatorWhitelist[operator];
    }

    function _setPoAManager(address _poaManager) internal {
        if (_poaManager == address(0)) {
            revert ZeroAddr();
        }
        address previousPoAManager = poaManager;
        poaManager = _poaManager;
        emit PoAManagerUpdated(previousPoAManager, _poaManager);
    }

    function _setCommitteeManager(address _committeeManager) internal {
        if (_committeeManager == address(0)) {
            revert ZeroAddr();
        }
        address previousCommitteeManager = committeeManager;
        committeeManager = _committeeManager;
        emit CommitteeManagerUpdated(previousCommitteeManager, _committeeManager);
    }

    function _setStateConfirmer(address _stateConfirmer) internal {
        if (_stateConfirmer == address(0)) {
            revert ZeroAddr();
        }
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
