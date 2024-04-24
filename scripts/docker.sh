#!/bin/bash -e

function build() {
    GIT_DATE=$(date +%FT%T%z)
    GIT_COMMIT=$(git rev-parse HEAD)
    docker-compose -f docker-compose-build.yaml build --build-arg BUILD_TAG=$BUILD_TAG --build-arg GIT_DATE=$GIT_DATE --build-arg GIT_COMMIT=$GIT_COMMIT "$@"
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