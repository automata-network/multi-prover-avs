pragma solidity ^0.8.12;

import "forge-std/Script.sol";

import "@dcap-v3-attestation/utils/SigVerifyLib.sol";
import "@dcap-v3-attestation/lib/PEMCertChainLib.sol";
import "@dcap-v3-attestation/AutomataDcapV3Attestation.sol";
import {TEELivenessVerifier} from "../src/core/TEELivenessVerifier.sol";

contract DeployTEELivenessVerifier is Script {
    function setUp() public {}

    function run() public {
        bool simulation = vm.envBool("SIMULATION");
        vm.startBroadcast();

        SigVerifyLib sigVerifyLib = new SigVerifyLib();
        PEMCertChainLib pemCertLib = new PEMCertChainLib();
        AutomataDcapV3Attestation attestation =
            new AutomataDcapV3Attestation(address(sigVerifyLib), address(pemCertLib));
        TEELivenessVerifier verifier = new TEELivenessVerifier(address(attestation), simulation);

        vm.stopBroadcast();

        string memory output = "tee liveness verifier contract";
        vm.serializeAddress(output, "SigVerifyLib", address(sigVerifyLib));
        vm.serializeAddress(output, "PEMCertChainLib", address(pemCertLib));
        vm.serializeAddress(output, "AutomataDcapV3Attestation", address(attestation));
        vm.serializeAddress(output, "TEELivenessVerifier", address(verifier));
        

        string memory outputFilePath = string.concat(vm.projectRoot(), "/script/output/tee_deploy_output.json");
        string memory finalJson = vm.serializeString(output, "object", output);
        vm.writeJson(finalJson, outputFilePath);
    }
}