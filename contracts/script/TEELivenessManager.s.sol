pragma solidity ^0.8.12;

import "forge-std/Script.sol";
import {TEELivenessVerifier} from "../src/core/TEELivenessVerifier.sol";

contract TEELivenessManager is Script {
    address serviceAddr = vm.envAddress("TEE_LIVENESS");
    TEELivenessVerifier liveness = TEELivenessVerifier(serviceAddr);

    function setUp() public {}

    function changeAttestValiditySeconds(uint256 secs) public {
        vm.startBroadcast();
        liveness.changeAttestValiditySeconds(secs);
    }

    function sendAttestation(TEELivenessVerifier.ReportDataV2 calldata reportData, bytes calldata quote) public {
        // string memory quote = vm.readFile(path);
        vm.startBroadcast();
        liveness.submitLivenessProofV2(reportData, quote);
    }
}
