#!/bin/bash -e

. $(dirname $0)/env.sh

function deployVerifier() {
    deploy deployVerifier $1 $2
}

function deployAll() {
    deploy all $1 $2
}

function deploy() {
    if [[ "$3" == "" ]]; then
        echo "usage: $0 $1 <proxy version> <attest validity secs>"
        return 1
    fi
    MAX_BLOCK_NUMBER_DIFF=$(_get_env "MAX_BLOCK_NUMBER_DIFF") \
    VERSION=$2 \
    ATTEST_VALIDITY_SECS=$3 \
    DEPLOY_KEY_SUFFIX=TEE_DEPLOY_KEY \
    ENV=$ENV \
    _script script/DeployTEELivenessService.s.sol --sig $1'()'

    cat $TEE_DEPLOY
}

function set_validity_secs() {
	TEE_LIVENESS=$(_get_key $TEE_DEPLOY .TEELivenessVerifierProxy) \
	_script script/TEELivenessManager.s.sol --sig 'changeAttestValiditySeconds(uint256)' $1
}


"$@"
