#!/bin/bash -e

function _aggregator_key() {
	_get_key config/aggregator.json .EcdsaPrivateKey
}

function _operator_key() {
	_get_key config/operator.json .EcdsaPrivateKey
}

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

function update_uri() {
	PRIVATE_KEY=$(_aggregator_key) \
	METADATA_URI=https://raw.githubusercontent.com/automata-network/multi-prover-avs-metadata/main/metadata.json \
	SERVICE_MANAGER_ADDRESS=$(_get_key $AVS_DEPLOY .multiProverServiceManager) \
	_script script/UpdateAVSMetadataURI.s.sol
}

function deploy_attestation() {
	SIMULATION=false
	if [[ "$1" == "--simulation" ]]; then
		SIMULATION=true
		echo "=============================================================="
		echo "NOTE: attestation contract are deployed with simulation mode!!"
		echo
	fi
	PRIVATE_KEY=$(_aggregator_key) SIMULATION=$SIMULATION _script script/DeployTEELivenessService.s.sol

	teeVerifierAddr=$(_get_key $TEE_DEPLOY .TEELivenessVerifier)
	_set_key config/aggregator.json TEELivenessVerifierContractAddress $teeVerifierAddr
	_set_key config/aggregator-docker-compose.json TEELivenessVerifierContractAddress $teeVerifierAddr
	_set_key config/operator.json TEELivenessVerifierAddr $teeVerifierAddr
	_set_key config/operator-docker-compose.json TEELivenessVerifierAddr $teeVerifierAddr

	if [[ "$1" == "--simulation" ]]; then
		echo
		echo "=============================================================="
		echo "NOTE: attestation contract are deployed with simulation mode!!"
		echo
	fi
}

function _script() {
	cd contracts
	forge script "$@" -v --broadcast --rpc-url $RPC_URL --private-key $PRIVATE_KEY
	cd -
}

function redeploy_avs() {
	topup_steth
	deploy_avs
}

function deploy_avs() {
	DELEGATION_MANAGER=$(_get_key $EIGENLAYER_DEPLOY .delegationManager) \
	AVS_DIRECTORY=$(_get_key $EIGENLAYER_DEPLOY .avsDirectory) \
	PRIVATE_KEY=$(_aggregator_key) _script script/DeployMultiProverServiceManager.s.sol

	registryCoordinator=$(_get_key $AVS_DEPLOY .registryCoordinator)
	_set_key config/operator.json RegistryCoordinatorAddress $registryCoordinator
	_set_key config/operator-docker-compose.json RegistryCoordinatorAddress $registryCoordinator
	_set_key config/aggregator.json AVSRegistryCoordinatorAddress $registryCoordinator
	_set_key config/aggregator-docker-compose.json AVSRegistryCoordinatorAddress $registryCoordinator

	operatorStateRetriever=$(_get_key $AVS_DEPLOY .operatorStateRetriever)
	_set_key config/aggregator.json OperatorStateRetrieverAddress $operatorStateRetriever
	_set_key config/aggregator-docker-compose.json OperatorStateRetrieverAddress $operatorStateRetriever

	multiProverServiceManager=$(_get_key $AVS_DEPLOY .multiProverServiceManager)
	_set_key config/aggregator.json MultiProverContractAddress $multiProverServiceManager
	_set_key config/aggregator-docker-compose.json MultiProverContractAddress $multiProverServiceManager
}

function topup_steth() {
	TARGET=$(_operator_key) \
	STRATEGY_ADDRESS=$(_get_key config/operator.json .strategyAddress) \
	_script script/StakeTokenTopup.s.sol
}

function topup_eth() {
	TARGET=$(_aggregator_key)
	if [[ "$1" == "operator" ]]; then
		TARGET=$(_operator_key)
	fi

	if [[ "$VALUE" == "" ]]; then
		export VALUE=200000000000000000
	fi

	TARGET=$TARGET VALUE=$VALUE _script script/EthTransfer.s.sol
}

function update_config() {
	_set_key config/aggregator.json EcdsaPrivateKey $AGGREGATOR_ECDSA_PRIVATE_KEY
	_set_key config/operator.json EcdsaPrivateKey $OPERATOR_ECDSA_PRIVATE_KEY
	_set_key config/operator.json BlsPrivateKey $OPERATOR_BLS_PRIVATE_KEY
	_set_key config/aggregator.json EthHttpEndpoint $RPC_URL
	_set_key config/aggregator.json EthWsEndpoint $RPC_WS
	_set_key config/operator.json EthRpcUrl $RPC_URL
	_set_key config/operator.json EthWsUrl $RPC_WS
}

function init_all() {
	# usage: 
	#   $0 init_all --simulation
	update_config
	topup_eth operator
	topup_eth aggregator
	topup_steth
	deploy_avs
	deploy_attestation $1 # enable simulation
}

source .env
if [[ "$PRIVATE_KEY" == "" ]]; then
	export PRIVATE_KEY=$TOPUP_PRIVATE_KEY
fi
AVS_DEPLOY=contracts/script/output/avs_deploy_output.json
TEE_DEPLOY=contracts/script/output/tee_deploy_output.json
EIGENLAYER_DEPLOY=contracts/script/output/eigenlayer_holesky_deploy.json


fn=$1
shift
$fn "$@"
