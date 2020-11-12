#!/usr/bin/env bash
set -ex

echo "This script automates the steps in the README"

bosh create-release --force

bosh upload-release

yes | bosh -d producer deploy manifests/lite_manifest.yml --no-redact
