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

import {IMultiProverServiceManager} from "../src/interfaces/IMultiProverServiceManager.sol";
import {MultiProverServiceManager} from "../src/core/MultiProverServiceManager.sol";

import "forge-std/Script.sol";

contract DeployMultiProverServiceManager is Script {
    function setUp() public {}

    function run() public {
         vm.startBroadcast();

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

        // Deploy and initialize MultiProverServiceManager
        IMultiProverServiceManager multiProverServiceManager;
        {
            multiProverServiceManager = IMultiProverServiceManager(
                address(new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), ""))
            );
            // Upgrade and initialize MultiProverServiceManager
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
                    msg.sender
                )
            );
        }
        
        // Deploy registryCoordinator implementation and initialize
        IRegistryCoordinator registryCoordinatorImpl = new RegistryCoordinator(multiProverServiceManager, stakeRegistry, blsApkRegistry, indexRegistry);              
        {
            // Upgrade and initialize RegistryCoordinator                        
            // 1 quorum, max 10 operator
            // 1.1 times the stake of original operators to kick it out
            // and the stake of original stackers should be less than 10% of the total stake
            IRegistryCoordinator.OperatorSetParam[] memory operatorSetParams = new IRegistryCoordinator.OperatorSetParam[](1);
            operatorSetParams[0] = IRegistryCoordinator.OperatorSetParam(uint32(10), uint16(11000), uint16(1000));
            uint96[] memory minimumStakeForQuourm = new uint96[](1);
            minimumStakeForQuourm[0] = uint96(1);
            IStakeRegistry.StrategyParams[][] memory strategyAndWeightingMultipliers = new IStakeRegistry.StrategyParams[][](1);
            {
                StrategyBaseTVLLimits deployedStrategy = StrategyBaseTVLLimits(vm.envAddress("STRATEGY"));
                strategyAndWeightingMultipliers[0] = new IStakeRegistry.StrategyParams[](1);
                strategyAndWeightingMultipliers[0][0] = IStakeRegistry.StrategyParams({
                    strategy: IStrategy(address(deployedStrategy)),
                    multiplier: 1 ether
                });
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

        vm.stopBroadcast();
    }
}