pragma solidity ^0.8.12;

import {IMultiProverServiceManager} from "../src/interfaces/IMultiProverServiceManager.sol";

import "forge-std/Script.sol";

contract TEEQuorumManagement is Script {
    IMultiProverServiceManager serviceManager = IMultiProverServiceManager(vm.envAddress("MULTI_PROVER_SERVICE_MANAGER"));
    uint256 privateKey = vm.envUint("PRIVATE_KEY");


    function addQuorum(IMultiProverServiceManager.TEE teeType, uint8 quorumNumber) public {
        vm.startBroadcast(privateKey);
        IMultiProverServiceManager.TEEQuorum memory quorum = IMultiProverServiceManager.TEEQuorum({
            teeType: teeType,
            quorumNumber: quorumNumber
        });
        serviceManager.addTEEQuorum(quorum);
        vm.stopBroadcast();
    }

    function removeQuorum(uint8 quorumNumber) public {
        vm.startBroadcast(privateKey);
        serviceManager.removeTEEQuorum(quorumNumber);
        vm.stopBroadcast();
    }
}