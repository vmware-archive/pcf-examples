#!/usr/bin/env bash
set -ex

echo "This script automates the steps in the README"

go_pkg_remote=https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz
go_pkg_path=./tmp/go-linux-amd64.tar.gz

if [ -f "${go_pkg_path}" ]; then
    echo "Go package already exist, skipping download"
else
    echo "Go package doesn't exist, downloading"
    wget "${go_pkg_remote}" -O "${go_pkg_path}"
fi
echo "${go_pkg_remote}" > ./tmp/go-version.txt

echo "Packaging local source"

tar -cvzf ./tmp/spacebears_src.tgz -C ../src/ spacebears/

echo "Adding blobs"

bosh add-blob ./tmp/go-linux-amd64.tar.gz go-linux-amd64.tar.gz
bosh add-blob ./tmp/go-version.txt go-version.txt
bosh add-blob ./tmp/spacebears_src.tgz spacebears_src.tgz

echo "Creating release"
bosh create-release --force --tarball ./tmp/

echo "Uploading release"
bosh upload-release

echo "Downloading routing release"
routing_release_remote=https://github.com/cloudfoundry-incubator/routing-release/releases/download/0.162.0/routing-0.162.0.tgz
routing_release_path=./tmp/routing-release.tgz
if [ ! -f "${routing_release_path}" ]; then
    wget "${routing_release_remote}" -O ${routing_release_path}
fi

echo "Uploading routing release"
bosh upload-release ${routing_release_path}

yes | bosh -d bosh_simple_with_routing deploy manifests/lite_manifest.yml --no-redact
