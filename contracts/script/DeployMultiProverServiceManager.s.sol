pragma solidity ^0.8.12;

import {TransparentUpgradeableProxy, ITransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {EmptyContract} from "eigenlayer-contracts/src/test/mocks/EmptyContract.sol";
import {IAVSDirectory} from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import {IPauserRegistry} from "eigenlayer-contracts/src/contracts/interfaces/IPauserRegistry.sol";
import {IDelegationManager} from "eigenlayer-contracts/src/contracts/interfaces/IDelegationManager.sol";
import {PauserRegistry} from "eigenlayer-contracts/src/contracts/permissions/PauserRegistry.sol";
import {StrategyBaseTVLLimits} from "eigenlayer-contracts/src/contracts/strategies/StrategyBaseTVLLimits.sol";

import {IStakeRegistry} from "eigenlayer-middleware/interfaces/IStakeRegistry.sol";
import {IIndexRegistry} from "eigenlayer-middleware/interfaces/IIndexRegistry.sol";
import {IBLSApkRegistry} from "eigenlayer-middleware/interfaces/IBLSApkRegistry.sol";
import {IRegistryCoordinator} from "eigenlayer-middleware/interfaces/IRegistryCoordinator.sol";
import {RegistryCoordinator} from "eigenlayer-middleware/RegistryCoordinator.sol";
import {IndexRegistry} from "eigenlayer-middleware/IndexRegistry.sol";
import {StakeRegistry, IStrategy} from "eigenlayer-middleware/StakeRegistry.sol";
import {BLSApkRegistry} from "eigenlayer-middleware/BLSApkRegistry.sol";
import {OperatorStateRetriever} from "eigenlayer-middleware/OperatorStateRetriever.sol";

import {IMultiProverServiceManager} from "../src/interfaces/IMultiProverServiceManager.sol";
import {MultiProverServiceManager} from "../src/core/MultiProverServiceManager.sol";

import "forge-std/Script.sol";

contract DeployMultiProverServiceManager is Script {
    function setUp() public {}

    function run() public {
        vm.startBroadcast();

        // These are deployed strategies used by EigenDA on Holesky testnet, we will also use them for multi-prover AVS
        // address[] memory quorum0Strategies = new address[](11);
        // {
        //     quorum0Strategies[0] = 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0; // Virtual strategy for beacon chain ETH
        //     quorum0Strategies[1] = 0x7D704507b76571a51d9caE8AdDAbBFd0ba0e63d3; // stETH strategy
        //     quorum0Strategies[2] = 0x3A8fBdf9e77DFc25d09741f51d3E181b25d0c4E0; // rETH strategy
        //     quorum0Strategies[3] = 0x05037A81BD7B4C9E0F7B430f1F2A22c31a2FD943; // lsETH strategy
        //     quorum0Strategies[4] = 0x9281ff96637710Cd9A5CAcce9c6FAD8C9F54631c; // sfrxETH strategy
        //     quorum0Strategies[5] = 0x31B6F59e1627cEfC9fA174aD03859fC337666af7; // ETHx strategy
        //     quorum0Strategies[6] = 0x46281E3B7fDcACdBa44CADf069a94a588Fd4C6Ef; // osETH strategy
        //     quorum0Strategies[7] = 0x70EB4D3c164a6B4A5f908D4FBb5a9cAfFb66bAB6; // cbETH strategy
        //     quorum0Strategies[8] = 0xaccc5A86732BE85b5012e8614AF237801636F8e5; // mETH strategy
        //     quorum0Strategies[9] = 0x7673a47463F80c6a3553Db9E54c8cDcd5313d0ac; // ankrETH strategy
        //     quorum0Strategies[10] = 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9; // WETH strategy
        // }

        // These are deployed strategies used by EigenDA on Ethereum mainnet, we will also use them for multi-prover AVS
        address[] memory quorum0Strategies = new address[](13);
        {
            quorum0Strategies[0] = 0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0; // Virtual strategy for beacon chain ETH
            quorum0Strategies[1] = 0x93c4b944D05dfe6df7645A86cd2206016c51564D; // stETH strategy
            quorum0Strategies[2] = 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2; // rETH strategy
            quorum0Strategies[3] = 0xAe60d8180437b5C34bB956822ac2710972584473; // lsETH strategy
            quorum0Strategies[4] = 0x8CA7A5d6f3acd3A7A8bC468a8CD0FB14B6BD28b6; // sfrxETH strategy
            quorum0Strategies[5] = 0x9d7eD45EE2E8FC5482fa2428f15C971e6369011d; // ETHx strategy
            quorum0Strategies[6] = 0x57ba429517c3473B6d34CA9aCd56c0e735b94c02; // osETH strategy
            quorum0Strategies[7] = 0x54945180dB7943c0ed0FEE7EdaB2Bd24620256bc; // cbETH strategy
            quorum0Strategies[8] = 0x298aFB19A105D59E74658C4C334Ff360BadE6dd2; // mETH strategy
            quorum0Strategies[9] = 0x13760F50a9d7377e4F20CB8CF9e4c26586c658ff; // ankrETH strategy
            quorum0Strategies[10] = 0xa4C637e0F704745D182e4D38cAb7E7485321d059; // OETH strategy
            quorum0Strategies[11] = 0x0Fe4F44beE93503346A3Ac9EE5A26b130a5796d6; // swETH strategy
            quorum0Strategies[12] = 0x7CA911E83dabf90C90dD3De5411a10F1A6112184; // wBETH strategy
        }

        EmptyContract emptyContract = new EmptyContract();
        ProxyAdmin proxyAdmin = new ProxyAdmin();

        IPauserRegistry pauserRegistry;
        {
            address[] memory pausers = new address[](1);
            pausers[0] = msg.sender;
            pauserRegistry = new PauserRegistry(pausers, msg.sender);
        }

        // Deploy indexRegistry, blsApkRegistry, stakeRegistry, registryCoordinator
        // For registryCoordinator, only proxy is deployed now
        IIndexRegistry indexRegistry;
        IBLSApkRegistry blsApkRegistry;
        IStakeRegistry stakeRegistry;
        IRegistryCoordinator registryCoordinator;
        {
            indexRegistry = IIndexRegistry(
                address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
            );
            blsApkRegistry = IBLSApkRegistry(
                address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
            );
            stakeRegistry = IStakeRegistry(
                address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
            );
            registryCoordinator = IRegistryCoordinator(
                address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
            );

            IIndexRegistry indexRegistryImpl = new IndexRegistry(registryCoordinator);
            IBLSApkRegistry blsApkRegistryImpl = new BLSApkRegistry(registryCoordinator);
            address delegationManager = vm.envAddress("DELEGATION_MANAGER");
            IStakeRegistry stakeRegistryImpl = new StakeRegistry(registryCoordinator, IDelegationManager(delegationManager));            

            proxyAdmin.upgrade(ITransparentUpgradeableProxy(address(indexRegistry)), address(indexRegistryImpl));
            proxyAdmin.upgrade(ITransparentUpgradeableProxy(address(blsApkRegistry)), address(blsApkRegistryImpl));
            proxyAdmin.upgrade(ITransparentUpgradeableProxy(address(stakeRegistry)), address(stakeRegistryImpl));            
        }        

        IMultiProverServiceManager multiProverServiceManager = IMultiProverServiceManager(
            address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
        );
        
        // Deploy registryCoordinator implementation and initialize
        IRegistryCoordinator registryCoordinatorImpl = new RegistryCoordinator(multiProverServiceManager, stakeRegistry, blsApkRegistry, indexRegistry);              
        {
            // Upgrade and initialize RegistryCoordinator
            // Quorum config:                      
            //      - max 50 operator
            //      - minimum stake is 0.01 ETH/LST
            //      - 1.1 times the stake of original operators to kick it out
            //      - the stake of original stackers should be less than 5% of the total stake
            IRegistryCoordinator.OperatorSetParam[] memory operatorSetParams = new IRegistryCoordinator.OperatorSetParam[](1);
            operatorSetParams[0] = IRegistryCoordinator.OperatorSetParam(uint32(50), uint16(11000), uint16(200));
            uint96[] memory minimumStakeForQuourm = new uint96[](1);
            minimumStakeForQuourm[0] = uint96(10000000000000000);
            IStakeRegistry.StrategyParams[][] memory strategyAndWeightingMultipliers = new IStakeRegistry.StrategyParams[][](1);
            {
                // address strategyManager = vm.envAddress("STRATEGY_MANAGER");
                // Same weight for all strategies
                strategyAndWeightingMultipliers[0] = new IStakeRegistry.StrategyParams[](quorum0Strategies.length);
                for (uint i = 0; i < quorum0Strategies.length; i++) {
                    // StrategyBaseTVLLimits strategy = StrategyBaseTVLLimits(strategyManager);
                    strategyAndWeightingMultipliers[0][i] = IStakeRegistry.StrategyParams({
                        strategy: IStrategy(quorum0Strategies[i]),
                        multiplier: 1 ether
                    });
                }
            }
            bytes memory callData;
            {
                callData = abi.encodeWithSelector(
                    RegistryCoordinator.initialize.selector,
                    msg.sender,
                    msg.sender,
                    msg.sender,
                    pauserRegistry,
                    0,
                    operatorSetParams,
                    minimumStakeForQuourm,
                    strategyAndWeightingMultipliers
                );
            }
            proxyAdmin.upgradeAndCall(
                ITransparentUpgradeableProxy(address(registryCoordinator)), 
                address(registryCoordinatorImpl),
                callData
            );
        }

        // Upgrade and initialize MultiProverServiceManager
        {
            address avsDirectory = vm.envAddress("AVS_DIRECTORY");
            IMultiProverServiceManager multiProverServiceManagerImpl = new MultiProverServiceManager(IAVSDirectory(avsDirectory), registryCoordinator, stakeRegistry);
            proxyAdmin.upgradeAndCall(
                ITransparentUpgradeableProxy(address(multiProverServiceManager)), 
                address(multiProverServiceManagerImpl),
                abi.encodeWithSelector(
                    MultiProverServiceManager.initialize.selector,
                    pauserRegistry,
                    0,
                    msg.sender,
                    msg.sender,
                    msg.sender,
                    msg.sender,
                    true    // Enable PoA
                )
            );
        }

        OperatorStateRetriever operatorStateRetriever = new OperatorStateRetriever();
        vm.stopBroadcast();

        string memory output = "multi-prover avs contracts deployment output";
        vm.serializeAddress(output, "indexRegistry", address(indexRegistry));
        vm.serializeAddress(output, "blsApkRegistry", address(blsApkRegistry));
        vm.serializeAddress(output, "stakeRegistry", address(stakeRegistry));
        vm.serializeAddress(output, "registryCoordinator", address(registryCoordinator));
        vm.serializeAddress(output, "multiProverServiceManager", address(multiProverServiceManager));
        vm.serializeAddress(output, "operatorStateRetriever", address(operatorStateRetriever));
        vm.serializeAddress(output, "proxyAdmin", address(proxyAdmin));

        string memory outputFilePath = string.concat(vm.projectRoot(), "/script/output/avs_deploy_output.json");
        string memory finalJson = vm.serializeString(output, "object", output);
        vm.writeJson(finalJson, outputFilePath);
    }
}