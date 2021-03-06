#!/bin/bash

set -euxo pipefail

# assume we are in the service directory
SERVICEDIR=$PWD
CLIENTDIR="${SERVICEDIR}/cmd/grpc_client"

(
	cd $CLIENTDIR
	go build -o=$GOPATH/bin/passwordservice_grpc_client
)
