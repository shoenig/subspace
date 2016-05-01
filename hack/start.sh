#!/bin/bash

set -u
set -x
set -e

old="$(pwd)"
function goback {
    cd ${old}
    echo ""
    echo ""
    echo "killing masters ${M1} ${M2} ${M3} ... "
    kill -9 ${M1} && kill -9 ${M2} && kill -9 ${M3}
    echo "killing agents ${A1} ${A2} ..."
    kill -9 ${A1} && kill -9 ${A2}
}
trap goback EXIT

cd "${GOPATH}/src/github.com/shoenig/subspace/cmd/subspace-master/"
go clean && go build
./subspace-master --bootstrap --config "${GOPATH}/src/github.com/shoenig/subspace/hack/master.config1.json" & export M1=$!
./subspace-master --config "${GOPATH}/src/github.com/shoenig/subspace/hack/master.config2.json" & export M2=$!
./subspace-master --config "${GOPATH}/src/github.com/shoenig/subspace/hack/master.config3.json" & export M3=$!
cd "${GOPATH}/src/github.com/shoenig/subspace/cmd/subspace-agent/"
go clean && go build
./subspace-agent --config "${GOPATH}/src/github.com/shoenig/subspace/hack/agent.config1.json" & export A1=$!
./subspace-agent --config "${GOPATH}/src/github.com/shoenig/subspace/hack/agent.config2.json" & export A2=$!

read nothing



