pragma solidity ^0.8.12;

import "forge-std/Script.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";

contract StackTokenTopup is Script {
    function setUp() public {}

    function run() public {
        address pk = vm.addr(uint256(vm.envBytes32("TARGET")));
        vm.startBroadcast();
        address strategyAddress = vm.envAddress("STRATEGY_ADDRESS");
        IStrategy strategy = IStrategy(strategyAddress);
        IERC20 token = strategy.underlyingToken();
        token.transfer(pk, 32000000000000000000);
        vm.stopBroadcast();
    }
}
