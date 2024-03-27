#!/bin/bash -e


function getKey() {
	file=$1
	key=$2
	cat $1 | jq $key | awk -F'"' '{print $2}'
}



function deploy_avs() {
	cd contracts
	# forge script script/EigenLayerDeployer.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast -vvvv

	DELEGATION_MANAGER=$(getKey script/output/eigenlayer_holesky_deploy.json .delegationManager) \
	AVS_DIRECTORY=$(getKey script/output/eigenlayer_holesky_deploy.json .avsDirectory) \
	forge script script/DeployMultiProverServiceManager.s.sol --rpc-url $RPC_URL --private-key $PRIVATE_KEY --broadcast -vvvv
}

function topup_steth() {
	privateKey=$1
	if [[ "$privateKey" == "" ]]; then
		echo "usage: $0 <private_key>" >&2
		exit 1
	fi
	strategyAddress=$(getKey config/operator.json .strategyAddress)
	cd contracts


	TARGET=$PRIVATE_KEY \
	STRATEGY_ADDRESS=$strategyAddress \
	forge script script/StakeTokenTopup.s.sol -vvvv --rpc-url $RPC_URL --private-key $privateKey --broadcast 
}


if [[ "$PRIVATE_KEY" == "" ]]; then
	export PRIVATE_KEY=$(getKey config/operator.json .EcdsaPrivateKey)
fi
RPC_URL=http://localhost:8545

fn=$1
shift
$fn "$@"