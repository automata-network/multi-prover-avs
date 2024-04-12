pragma solidity ^0.8.12;

import {MultiProverServiceManager} from "../src/core/MultiProverServiceManager.sol";
import {IMultiProverServiceManager} from "../src/interfaces/IMultiProverServiceManager.sol";

import "forge-std/Script.sol";

contract Whitelist is Script {
    address serviceAddr = vm.envAddress("MULTI_PROVER_SERVICE_MANAGER");
    MultiProverServiceManager serviceManager = MultiProverServiceManager(serviceAddr);


    function add(address operator) public {
        vm.startBroadcast();
        serviceManager.whitelistOperator(operator);
        vm.stopBroadcast();
    }

    function remove(address operator) public {
        vm.startBroadcast();
        serviceManager.blacklistOperator(operator);
        vm.stopBroadcast();
    }

    function addQuorum(IMultiProverServiceManager.TEE teeType, uint8 quorumNumber) public {
        vm.startBroadcast();
        IMultiProverServiceManager.TEEQuorum memory quorum = IMultiProverServiceManager.TEEQuorum({
            teeType: teeType,
            quorumNumber: quorumNumber
        });
        serviceManager.addTEEQuorum(quorum);
        vm.stopBroadcast();
    }

    function addCommittee(uint256 id, string memory description, bytes memory metadata, bytes memory teeQuorumNumbers) public {
        vm.startBroadcast();
        IMultiProverServiceManager.Committee memory committee = IMultiProverServiceManager.Committee({
            id: id,
            description: description,
            metadata: metadata,
            teeQuorumNumbers: teeQuorumNumbers
        });
        serviceManager.addCommittee(committee);
        vm.stopBroadcast();
    }
}