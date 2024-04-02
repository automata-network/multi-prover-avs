pragma solidity ^0.8.12;

import {IMultiProverServiceManager} from "../src/interfaces/IMultiProverServiceManager.sol";

import "forge-std/Script.sol";

contract UpdateAVSMetadataURI is Script {
    function setUp() public {}

    function run() public {
        vm.startBroadcast();
        address avsAddress = vm.envAddress("SERVICE_MANAGER_ADDRESS");
        string memory metadataURI = vm.envString("METADATA_URI");
        IMultiProverServiceManager avs = IMultiProverServiceManager(avsAddress);
        avs.updateAVSMetadataURI(metadataURI);
        vm.stopBroadcast();
    }
}
