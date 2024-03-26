# Build with Foundry
Compile the contract:
```
forge build
```
Testing the contract:
```
forge test
```
Deploy Eigenlayer core contracts(can be skipped if you want to use contracts deployed by Eigenlayer):
```
forge script script/EigenLayerDeployer.s.sol --rpc-url <rpc-url> --private-key <deployer-private-key> --broadcast -vvvv
```
Deploy Multi-prover AVS contracts:
Env can be found inside `./script/output/eigenlayer_deploy_output.json` if you want to use the eigenlayer contracts deployed at previous step, or get the deployed address from [Eigenlayer's docs](https://docs.eigenlayer.xyz/eigenlayer/deployed-contracts/).
```
DELEGATION_MANAGER=<delegation-manager-address> AVS_DIRECTORY=<avs-directory-address> forge script script/DeployMultiProverServiceManager.s.sol --rpc-url <rpc-url> --private-key <deployer-private-key> --broadcast -vvvv
```