#!/usr/bin/env bash
set -x

pushd ../spacebears

# go compile commands
GOOS=linux GOARCH=amd64 go build main.go
mv main ../db-docker/spacebears

popd

# some docker commands
docker build . -t cfplatformeng/spacebears
