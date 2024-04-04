#!/bin/bash -e

function _get_key() {
	file=$1
	key=$2
	cat $1 | jq $key | awk -F'"' '{print $2}'
}

function _set_key() {
	file=$1
	key=$2
	value=$3
	if [[ $(grep "\\\"$key\\\"" $file) == "" ]]; then
		echo "key $key not found in $file"
		return 1
	fi
	sed "s!\\\"$key\\\": \\\".*\\\"!\\\"$key\": \\\"$value\\\"!g" $file > /tmp/tmp_set_key
	cat /tmp/tmp_set_key > $file
}


function update_config() {
	_set_key config/aggregator.json EcdsaPrivateKey $AGGREGATOR_ECDSA_PRIVATE_KEY
        _set_key config/operator.json EcdsaPrivateKey $OPERATOR_ECDSA_PRIVATE_KEY
	_set_key config/operator.json BlsPrivateKey $OPERATOR_BLS_PRIVATE_KEY
	_set_key config/operator.json strategyAddress $STRATEGY_ADDRESS

	_set_key config/aggregator.json EthHttpEndpoint $RPC_URL
	_set_key config/aggregator.json EthWsEndpoint $RPC_WS
	_set_key config/operator.json EthRpcUrl $RPC_URL
	_set_key config/operator.json EthWsUrl $RPC_WS

	_set_key config/aggregator.json Simulation $SIMULATION
	_set_key config/operator.json simulation $SIMULATION
	
	_set_key config/operator.json RegistryCoordinatorAddress $REGISTRY_COORDINATOR_ADDRESS
	_set_key config/aggregator.json AVSRegistryCoordinatorAddress $REGISTRY_COORDINATOR_ADDRESS
	
	_set_key config/aggregator.json OperatorStateRetrieverAddress $OPERATOR_STATE_RETRIEVER_ADDRESS
	
	_set_key config/aggregator.json MultiProverContractAddress $MULTI_PROVER_CONTRACT_ADDRESS
	
	_set_key config/aggregator.json TEELivenessVerifierContractAddress $TEE_LIVENESS_VERIFIER_ADDRESS
	_set_key config/operator.json TEELivenessVerifierAddr $TEE_LIVENESS_VERIFIER_ADDRESS
}

function run() {
	docker-compose -f docker-compose-holesky.yaml up
}

function stop() {
    docker-compose down
}

source .env
if [[ "$PRIVATE_KEY" == "" ]]; then
	export PRIVATE_KEY=$TOPUP_PRIVATE_KEY
fi

fn=$1
shift
$fn "$@"
