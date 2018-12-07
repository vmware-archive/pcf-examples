#!/usr/bin/env bash
set -e

echo "This script builds a tile that packages a simple addon"

echo "Downloading release"
curl -L https://bosh.io/d/github.com/cloudfoundry/os-conf-release?v=20.0.0 > resources/os-conf-release.20.0.0.tgz

tile build
