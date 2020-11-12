#!/usr/bin/env bash
set -ex

bosh create-release --tarball producer.tgz --force
mv producer.tgz resources

tile build
