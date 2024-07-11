#!/bin/bash -e

. $(dirname $0)/env.sh

function add_whitelist() {
    MULTI_PROVER_SERVICE_MANAGER=$(_get_key $AVS_DEPLOY .multiProverServiceManager) \
    DEPLOY_KEY_SUFFIX=DEPLOY_KEY \
    _script script/Whitelist.s.sol --sig 'add(address)' $1
}

function changeMaxBlockNumberDiff() {
    TEE_LIVENESS=$(_get_key $TEE_DEPLOY .TEELivenessVerifierProxy) \
    DEPLOY_KEY_SUFFIX=DEPLOY_KEY \
    _script script/TEELivenessManager.s.sol --sig 'changeMaxBlockNumberDiff(uint256)' $1
}

function addLineaQuorum() {
    MULTI_PROVER_SERVICE_MANAGER=$(_get_key $AVS_DEPLOY .multiProverServiceManager) \
    DEPLOY_KEY_SUFFIX=AVS_DEPLOY_KEY \
    _script script/TEECommitteeManagement.s.sol --sig 'addLineaQuorum()'
}

function addLineaCommittee() {
    MULTI_PROVER_SERVICE_MANAGER=$(_get_key $AVS_DEPLOY .multiProverServiceManager) \
    DEPLOY_KEY_SUFFIX=AVS_DEPLOY_KEY \
    _script script/TEECommitteeManagement.s.sol --sig 'addLineaCommittee()'
}

"$@"