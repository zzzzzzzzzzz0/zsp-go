#!/bin/sh
if [ $GOPATH ]; then
	GOPATH=:$GOPATH
fi
export GOPATH=$PWD:/opt2/go$GOPATH
/usr/bin/go install -v -gcflags "-N -l" ./...
