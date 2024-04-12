//SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {IMultiProverServiceManager} from "../interfaces/IMultiProverServiceManager.sol";

abstract contract MultiProverServiceManagerStorage is IMultiProverServiceManager {
    /// @notice The pause flag for submitting state transitions
    uint8 public constant PAUSED_SUBMIT_STATE = 0;
    /// @notice The pause flag for operator registration
    uint8 public constant PAUSED_OPERTOR_REGISTRATION = 1;

    uint256 public constant THRESHOLD_DENOMINATOR = 100;
    
    /// @notice The current task ID
    uint32 public taskId;
    
    /// @notice The task metadata hash for a given task ID
    mapping(uint32 => bytes32) public taskIdToMetadataHash;

    mapping(uint256 => Committee) public committees;

    mapping(uint8 => TEEQuorum) public teeQuorums;

    mapping(uint8 => uint256) public quorumIdToCommitteeId;

    /// @notice the address that is permissioned to submit state transitions
    address public stateConfirmer;

    address public poaManager;

    address public committeeManager;

    mapping(address => bool) internal operatorWhitelist;

    bool public poaEnabled;
}