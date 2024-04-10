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
        require(msg.sender == stateConfirmer, "MultiProverServiceManager.onlyStateConfirmer: caller is not the state confirmer");
        _;
    }

    modifier onlyPoAManager() {
        require(msg.sender == poaManager, "MultiProverServiceManager.onlyPoAManager: caller is not the PoA manager");
        _;
    }

    modifier onlyCommitteeManager() {
        require(msg.sender == committeeManager, "MultiProverServiceManager.onlyCommitteeManager: caller is not the committee manager");
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
        require(!poaEnabled || operatorWhitelist.contains(operator), "MultiProverServiceManager.registerOperatorToAVS: operator not whitelisted");
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
        require(tx.origin == msg.sender, "MultiProverServiceManager.confirmState: state transition and nonsigner data must be in calldata");

        //make sure that the quorumNumbers and signedStakeForQuorums are of the same length
        require(
            stateHeader.quorumNumbers.length == stateHeader.quorumThresholdPercentages.length,
            "MultiProverServiceManager.confirmState: quorumNumbers and signedStakeForQuorums must be of the same length"
        );

        // make sure the committee exists
        require(committees[stateHeader.committeeId].id != 0, "MultiProverServiceManager.confirmState: committee does not exist");

        // make sure the quorums belong to the committee
        for (uint256 i = 0; i < stateHeader.quorumNumbers.length; i++) {
            require(
                quorumIdToCommitteeId[uint8(stateHeader.quorumNumbers[i])] == stateHeader.committeeId,
                "MultiProverServiceManager.confirmState: quorum does not belong to committee"
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
                "MultiProverServiceManager.confirmState: signatories do not own at least threshold percentage of a quorum"
            );
        }

        uint32 taskIdMemory = taskId;
        bytes32 stateHeaderHash = _hashStateHeader(stateHeader);
        bytes32 taskMetadataHash = keccak256(abi.encodePacked(stateHeaderHash, signatoryRecordHash, uint32(block.number)));
        taskIdToMetadataHash[taskIdMemory] = taskMetadataHash;

        taskId = taskIdMemory + 1;

        emit StateConfirmed(stateHeader.committeeId, stateHeader.metadata, stateHeader.state);
    }

    function addCommittee(Committee memory committee) external onlyCommitteeManager {
        require(committee.id != 0, "MultiProverServiceManager.addCommittee: committee id cannot be 0");
        require(committees[committee.id].id == 0, "MultiProverServiceManager.addCommittee: committee already exists");
        for (uint256 i = 0; i < committee.teeQuorumNumbers.length; i++) {
            uint8 teeQuorumNumber = uint8(committee.teeQuorumNumbers[i]);
            require(teeQuorums[teeQuorumNumber].teeType != TEE.NONE, "MultiProverServiceManager.addCommittee: tee quorum does not exist");
            require(quorumIdToCommitteeId[teeQuorumNumber] == 0, "MultiProverServiceManager.addCommittee: tee quorum is used by another committee");
            quorumIdToCommitteeId[teeQuorumNumber] = committee.id;
        }

        committees[committee.id] = committee;
    }

    function updateCommittee(Committee memory committee) external onlyCommitteeManager {
        require(committees[committee.id].id != 0, "MultiProverServiceManager.updateCommittee: committee does not exist");

        _removeCommittee(committee.id);
        for (uint256 i = 0; i < committee.teeQuorumNumbers.length; i++) {
            uint8 teeQuorumNumber = uint8(committee.teeQuorumNumbers[i]);
            require(teeQuorums[teeQuorumNumber].teeType != TEE.NONE, "MultiProverServiceManager.addCommittee: tee quorum does not exist");
            require(quorumIdToCommitteeId[teeQuorumNumber] == 0, "MultiProverServiceManager.updateCommittee: tee quorum is used by another committee");
            quorumIdToCommitteeId[teeQuorumNumber] = committee.id;
        }

        committees[committee.id] = committee;
    }

    function removeCommittee(uint256 committeeId) external onlyCommitteeManager {
        _removeCommittee(committeeId);
    }

    function _removeCommittee(uint256 committeeId) internal {
        require(committees[committeeId].id != 0, "MultiProverServiceManager.removeCommittee: committee does not exist");
        bytes memory teeQuorumNumbers = committees[committeeId].teeQuorumNumbers;
        for (uint256 i = 0; i < teeQuorumNumbers.length; i++) {
            quorumIdToCommitteeId[uint8(teeQuorumNumbers[i])] = 0;
        }

        delete committees[committeeId];
    }

    function addTEEQuorum(TEEQuorum memory teeQuorum) external onlyCommitteeManager {
        require(teeQuorums[teeQuorum.quorumNumber].teeType == TEE.NONE, "MultiProverServiceManager.addTEEQuorum: tee quorum already exists");
        require(teeQuorum.teeType != TEE.NONE, "MultiProverServiceManager.addTEEQuorum: tee type cannot be NONE");
        require(_stakeRegistry.getTotalStakeHistoryLength(teeQuorum.quorumNumber) != 0, "MultiProverServiceManager.addTEEQuorum: quorum not initialized");

        teeQuorums[teeQuorum.quorumNumber] = teeQuorum;
    }

    function removeTEEQuorum(uint8 quorumNumber) external onlyCommitteeManager {
        require(teeQuorums[quorumNumber].teeType != TEE.NONE, "MultiProverServiceManager.removeTEEQuorum: tee quorum does not exist");
        require(quorumIdToCommitteeId[quorumNumber] == 0, "MultiProverServiceManager.removeTEEQuorum: tee quorum is in use");

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
        require(!poaEnabled, "MultiProverServiceManager.enablePoA: PoA already enabled");
        poaEnabled = true;
    }

    function disablePoA() external onlyPoAManager {
        require(poaEnabled, "MultiProverServiceManager.disablePoA: PoA already disabled");
        poaEnabled = false;
    }

    function isPoAEnabled() external view returns (bool) {
        return poaEnabled;
    }

    function whitelistOperator(address operator) external onlyPoAManager {
        require(operator != address(0), "MultiProverServiceManager.whitelistOperator: operator cannot be the zero address");
        require(!operatorWhitelist.contains(operator), "MultiProverServiceManager.whitelistOperator: operator already whitelisted");
        operatorWhitelist.add(operator);
    }

    function blacklistOperator(address operator) external onlyPoAManager {
        require(operator != address(0), "MultiProverServiceManager.blacklistOperator: operator cannot be the zero address");
        require(operatorWhitelist.contains(operator), "MultiProverServiceManager.blacklistOperator: operator not whitelisted");
        operatorWhitelist.remove(operator);
    }

    function isOperatorWhitelisted(address operator) external view returns (bool) {
        return operatorWhitelist.contains(operator);
    }

    function _setPoAManager(address _poaManager) internal {
        require(_poaManager != address(0), "MultiProverServiceManager._setPoAManager: PoA manager cannot be the zero address");
        address previousPoAManager = poaManager;
        poaManager = _poaManager;
        emit PoAManagerUpdated(previousPoAManager, _poaManager);
    }

    function _setCommitteeManager(address _committeeManager) internal {
        require(_committeeManager != address(0), "MultiProverServiceManager._setCommitteeManager: committee manager cannot be the zero address");
        address previousCommitteeManager = committeeManager;
        committeeManager = _committeeManager;
        emit CommitteeManagerUpdated(previousCommitteeManager, _committeeManager);
    }

    function _setStateConfirmer(address _stateConfirmer) internal {
        require(_stateConfirmer != address(0), "MultiProverServiceManager._setStateConfirmer: state confirmer cannot be the zero address");
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