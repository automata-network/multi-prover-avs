pragma solidity ^0.8.12;

import {IMultiProverServiceManager} from "../src/interfaces/IMultiProverServiceManager.sol";

import "forge-std/Script.sol";

contract TEECommitteeManagement is Script {
    IMultiProverServiceManager serviceManager = IMultiProverServiceManager(vm.envAddress("MULTI_PROVER_SERVICE_MANAGER"));
    
    function run() public {
        uint256 id = 1;
        string memory description = "Scroll Prover Committee";
        bytes memory metadata = abi.encodePacked('{"chainId":534352}');
        bytes memory teeQuorumNumbers = new bytes(1);
        teeQuorumNumbers[0] = bytes1(uint8(0));
        addCommittee(id, description, metadata, teeQuorumNumbers);
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

    function updateCommittee(uint256 id, string memory description, bytes memory metadata, bytes memory teeQuorumNumbers) public {
        vm.startBroadcast();
        IMultiProverServiceManager.Committee memory committee = IMultiProverServiceManager.Committee({
            id: id,
            description: description,
            metadata: metadata,
            teeQuorumNumbers: teeQuorumNumbers
        });
        serviceManager.updateCommittee(committee);
        vm.stopBroadcast();
    }

    function removeCommittee(uint256 committeeId) public {
        vm.startBroadcast();
        serviceManager.removeCommittee(committeeId);
        vm.stopBroadcast();
    }
}