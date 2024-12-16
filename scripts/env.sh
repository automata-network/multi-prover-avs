#!/bin/bash

function _script() {
	if [[ "$WORKDIR" == "" ]]; then
		WORKDIR=contracts
	fi
	if [[ "$NOSEND" == "" ]]; then
		NOSEND="--broadcast"
	else
		NOSEND=""
	fi
	cd $WORKDIR
	# ETH_GAS_PRICE=0.1gwei
	ENV=$ENV forge script "$@" -vvv $NOSEND --rpc-url $(_get_env RPC_URL) --private-key $(_get_env $DEPLOY_KEY_SUFFIX) $(_get_env VERIFY_CMD)
	cd -
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
	sed "s/\\\"$key\\\": \\\".*\\\"/\\\"$key\": \\\"$value\\\"/g" $file > /tmp/tmp_set_key
	cat /tmp/tmp_set_key > $file
}

function _get_env() {
    env_name=$(echo "$ENV" | tr '[:lower:]' '[:upper:]')
    if [[ "$env_name" == "" ]]; then
        eval "echo \$${env_name}$1"
    else
        eval "echo \$${env_name}_$1"
    fi
}

# init env
. .env
if [[ "$ENV" == "" ]]; then
    ENV="localhost"
fi
echo "=============================================================="
echo "ENV: $ENV"
echo "RPC_URL: $(_get_env RPC_URL)"
echo "=============================================================="
echo

AVS_DEPLOY=contracts/script/output/avs_deploy_$ENV.json
TEE_DEPLOY=contracts/script/output/tee_deploy_output_$ENV.json
EIGENLAYER_DEPLOY=contracts/script/output/eigenlayer_holesky_deploy.json
PRIVATE_KEY=$(_get_env DEPLOY_KEY)

