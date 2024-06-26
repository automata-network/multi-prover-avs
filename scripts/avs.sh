#!/bin/bash -e

. $(dirname $0)/env.sh

function add_whitelist() {
    MULTI_PROVER_SERVICE_MANAGER=$(_get_key $AVS_DEPLOY .multiProverServiceManager) \
    DEPLOY_KEY_SUFFIX=DEPLOY_KEY \
    _script script/Whitelist.s.sol --sig 'add(address)' $1
}

"$@"