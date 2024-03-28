pragma solidity ^0.8.12;

import {IMultiProverServiceManager} from "../src/interfaces/IMultiProverServiceManager.sol";

import "forge-std/Script.sol";

contract TEECommitteeManagement is Script {
    IMultiProverServiceManager serviceManager = IMultiProverServiceManager(vm.envAddress("MULTI_PROVER_SERVICE_MANAGER"));
    uint256 privateKey = vm.envUint("PRIVATE_KEY");

    function addCommittee(uint256 id, string memory description, bytes memory metadata, bytes memory teeQuorumNumbers) public {
        vm.startBroadcast(privateKey);
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
        vm.startBroadcast(privateKey);
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
        vm.startBroadcast(privateKey);
        serviceManager.removeCommittee(committeeId);
        vm.stopBroadcast();
    }
}