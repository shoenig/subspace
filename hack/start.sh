#!/bin/bash

# use from cmd/subspace-{master,agent}
# for example,
#
# $ ../../hack/start.sh 1
# in subspace-master/ will start the first master.

set -e
set -u

go clean && go build

dir=${PWD##*/}

if [[ "${dir}" == "subspace-master" ]] ; then
    if [[ "${1}" == "1" ]] ; then
        bootstrap="--bootstrap"
    else
        bootstrap=""
    fi
    exec ./subspace-master "${bootstrap}" --config "../../hack/master.config${1}.json"
elif [[ "${dir}" == "subspace-agent" ]] ; then
    exec ./subspace-agent --config "../../hack/agent.config${1}.json"
else
    echo "what are you doing"
    exit 1
fi