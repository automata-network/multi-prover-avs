pragma solidity ^0.8.12;

import {IMultiProverServiceManager} from "../src/interfaces/IMultiProverServiceManager.sol";

import "forge-std/Script.sol";

contract TEEQuorumManagement is Script {

    IMultiProverServiceManager serviceManager =
        IMultiProverServiceManager(vm.envAddress("MULTI_PROVER_SERVICE_MANAGER"));

    function run() public {
        addQuorum(IMultiProverServiceManager.TEE.INTEL_SGX, 0);
    }

    function addQuorum(IMultiProverServiceManager.TEE teeType, uint8 quorumNumber) public {
        vm.startBroadcast();
        IMultiProverServiceManager.TEEQuorum memory quorum =
            IMultiProverServiceManager.TEEQuorum({teeType: teeType, quorumNumber: quorumNumber});
        serviceManager.addTEEQuorum(quorum);
        vm.stopBroadcast();
    }

    function removeQuorum(uint8 quorumNumber) public {
        vm.startBroadcast();
        serviceManager.removeTEEQuorum(quorumNumber);
        vm.stopBroadcast();
    }

}
