#!/usr/bin/env bash
set -ex

bosh create-release --tarball consumer.tgz --force
mv consumer.tgz resources

tile build
