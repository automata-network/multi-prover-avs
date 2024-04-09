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
	_set_key config/operator.json ProverURL $PROVER_URL
	
	_set_key config/operator.json RegistryCoordinatorAddress $REGISTRY_COORDINATOR_ADDRESS
	_set_key config/operator.json TEELivenessVerifierAddr $TEE_LIVENESS_VERIFIER_ADDRESS
	
	_set_bool_key config/operator.json simulation $SIMULATION

}

function topup_operator() {
	topup_eth
	topup_steth
}

function topup_steth() {
	TARGET=$OPERATOR_ECDSA_PRIVATE_KEY \
	STRATEGY_ADDRESS=$STRATEGY_ADDRESS \
	_script script/StakeTokenTopup.s.sol
}

function topup_eth() {
	if [[ "$VALUE" == "" ]]; then
		export VALUE=200000000000000000
	fi

	TARGET=$OPERATOR_ECDSA_PRIVATE_KEY VALUE=$VALUE _script script/EthTransfer.s.sol
}

function _script() {
	cd contracts
	forge script "$@" -v --broadcast --rpc-url $RPC_URL --private-key $TOPUP_PRIVATE_KEY
	cd -
}

function run() {
	docker-compose -f docker-compose-operator.yaml up
}

function stop() {
    docker-compose down
}

source .env

fn=$1
shift
$fn "$@"
