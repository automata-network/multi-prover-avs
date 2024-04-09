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

function _set_bool_key() {
        file=$1
        key=$2
        value=$3
        if [[ $(grep "\\\"$key\\\"" $file) == "" ]]; then
                echo "key $key not found in $file"
                return 1
        fi
        sed "s!\"$key\": \(true\|false\)!\"$key\": $value!g" $file > /tmp/tmp_set_key
        cat /tmp/tmp_set_key > $file
}


function update_config() {
        _set_key config/operator.json EcdsaPrivateKey $OPERATOR_ECDSA_PRIVATE_KEY
	_set_key config/operator.json BlsPrivateKey $OPERATOR_BLS_PRIVATE_KEY
	_set_key config/operator.json strategyAddress $STRATEGY_ADDRESS

	_set_key config/operator.json EthRpcUrl $RPC_URL
	_set_key config/operator.json EthWsUrl $RPC_WS
	
	_set_key config/operator.json AggregatorURL $AGGREGATOR_URL
	_set_key config/operator.json RegistryCoordinatorAddress $REGISTRY_COORDINATOR_ADDRESS
	_set_key config/operator.json TEELivenessVerifierAddr $TEE_LIVENESS_VERIFIER_ADDRESS
}

function run() {
	docker-compose -f docker-compose-holesky.yaml up
}

function stop() {
    docker-compose down
}

source .env

fn=$1
shift
$fn "$@"
