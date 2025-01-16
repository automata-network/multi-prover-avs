pragma solidity ^0.8.12;

import "forge-std/Script.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";

contract EthTransfer is Script {

    function setUp() public {}

    function run() public {
        address pk = vm.addr(uint256(vm.envBytes32("TARGET")));
        uint256 value = vm.envUint("VALUE");

        vm.startBroadcast();
        (bool success,) = pk.call{value: value}(new bytes(0));
        require(success);
        vm.stopBroadcast();
    }

}
