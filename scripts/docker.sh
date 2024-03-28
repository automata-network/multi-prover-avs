#!/bin/bash -e

function build() {
    docker-compose -f docker-compose-build.yaml build
}

function stop() {
    docker-compose down
    docker-compose -f docker-compose-state.yaml down
}

function run() {
    docker-compose up
}

function init_state() {
	stop
	docker compose -f docker-compose-state.yaml up --wait
	$(dirname $0)/deploy.sh init_all $@
}

cmd=$1
shift
$cmd "$@"