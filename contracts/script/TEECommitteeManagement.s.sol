pragma solidity ^0.8.12;

import {IMultiProverServiceManager} from "../src/interfaces/IMultiProverServiceManager.sol";
import {IStakeRegistry} from "eigenlayer-middleware/interfaces/IStakeRegistry.sol";
import {StakeRegistry, IStrategy} from "eigenlayer-middleware/StakeRegistry.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/interfaces/IRegistryCoordinator.sol";
import {RegistryCoordinator} from "eigenlayer-middleware/RegistryCoordinator.sol";

import "forge-std/Script.sol";

contract TEECommitteeManagement is Script {
    IMultiProverServiceManager serviceManager =
        IMultiProverServiceManager(
            vm.envAddress("MULTI_PROVER_SERVICE_MANAGER")
        );

    function run() public {
        uint256 id = 1;
        string memory description = "Scroll Prover Committee";
        bytes memory metadata = abi.encodePacked('{"chainId":534352}');
        bytes memory teeQuorumNumbers = new bytes(1);
        teeQuorumNumbers[0] = bytes1(uint8(0));
        addCommittee(id, description, metadata, teeQuorumNumbers);
    }

    function addLineaCommittee() public {
        uint256 id = 2;
        string memory description = "Linea Prover Committee";
        bytes memory metadata = abi.encodePacked('{"chainId":59144}');
        bytes memory teeQuorumNumbers = new bytes(1);
        teeQuorumNumbers[0] = bytes1(uint8(1));


        vm.startBroadcast();
        IMultiProverServiceManager.TEEQuorum memory teeQuorum = IMultiProverServiceManager.TEEQuorum({
            teeType: IMultiProverServiceManager.TEE.INTEL_SGX,
            quorumNumber: 1
        });
        serviceManager.addTEEQuorum(teeQuorum);
        vm.stopBroadcast();
        
        addCommittee(id, description, metadata, teeQuorumNumbers);
    }

    function addCommittee(
        uint256 id,
        string memory description,
        bytes memory metadata,
        bytes memory teeQuorumNumbers
    ) public {
        vm.startBroadcast();
        IMultiProverServiceManager.Committee
            memory committee = IMultiProverServiceManager.Committee({
                id: id,
                description: description,
                metadata: metadata,
                teeQuorumNumbers: teeQuorumNumbers
            });
        serviceManager.addCommittee(committee);
        vm.stopBroadcast();
    }

    function updateCommittee(
        uint256 id,
        string memory description,
        bytes memory metadata,
        bytes memory teeQuorumNumbers
    ) public {
        vm.startBroadcast();
        IMultiProverServiceManager.Committee
            memory committee = IMultiProverServiceManager.Committee({
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

    function addLineaQuorum() public {
        vm.startBroadcast();
        // holesky
        address[] memory strategies = new address[](11);
        {
            strategies[0] = 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0; // Virtual strategy for beacon chain ETH
            strategies[1] = 0x7D704507b76571a51d9caE8AdDAbBFd0ba0e63d3; // stETH strategy
            strategies[2] = 0x3A8fBdf9e77DFc25d09741f51d3E181b25d0c4E0; // rETH strategy
            strategies[3] = 0x05037A81BD7B4C9E0F7B430f1F2A22c31a2FD943; // lsETH strategy
            strategies[4] = 0x9281ff96637710Cd9A5CAcce9c6FAD8C9F54631c; // sfrxETH strategy
            strategies[5] = 0x31B6F59e1627cEfC9fA174aD03859fC337666af7; // ETHx strategy
            strategies[6] = 0x46281E3B7fDcACdBa44CADf069a94a588Fd4C6Ef; // osETH strategy
            strategies[7] = 0x70EB4D3c164a6B4A5f908D4FBb5a9cAfFb66bAB6; // cbETH strategy
            strategies[8] = 0xaccc5A86732BE85b5012e8614AF237801636F8e5; // mETH strategy
            strategies[9] = 0x7673a47463F80c6a3553Db9E54c8cDcd5313d0ac; // ankrETH strategy
            strategies[10] = 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9; // WETH strategy
        }
        IStakeRegistry.StrategyParams[]
            memory strategyParams = new IStakeRegistry.StrategyParams[](11);
        {
            for (uint i = 0; i < strategies.length; i++) {
                strategyParams[i] = IStakeRegistry.StrategyParams({
                    strategy: IStrategy(strategies[i]),
                    multiplier: 1 ether
                });
            }
        }

        uint96 minimumStakeForQuourm = 10000000000000000;
        IRegistryCoordinator.OperatorSetParam
            memory operatorSetParams = IRegistryCoordinator.OperatorSetParam(
                uint32(100),
                uint16(11000),
                uint16(100)
            );

        string memory output = readJson();
        RegistryCoordinator registryCoordinator = RegistryCoordinator(
            vm.parseJsonAddress(output, ".registryCoordinator")
        );

        registryCoordinator.createQuorum(
            operatorSetParams,
            minimumStakeForQuourm,
            strategyParams
        );
        vm.stopBroadcast();
    }

    function getOutputFilePath() private view returns (string memory) {
        string memory env = vm.envString("ENV");
        return
            string.concat(
                vm.projectRoot(),
                "/script/output/avs_deploy_",
                env,
                ".json"
            );
    }

    function readJson() private returns (string memory) {
        bytes32 remark = keccak256(abi.encodePacked("remark"));
        string memory output = vm.readFile(getOutputFilePath());
        string[] memory keys = vm.parseJsonKeys(output, ".");
        for (uint i = 0; i < keys.length; i++) {
            if (keccak256(abi.encodePacked(keys[i])) == remark) {
                continue;
            }
            string memory keyPath = string(abi.encodePacked(".", keys[i]));
            vm.serializeAddress(
                output,
                keys[i],
                vm.parseJsonAddress(output, keyPath)
            );
        }
        return output;
    }
}
