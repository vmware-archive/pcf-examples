#!/usr/bin/env bash
set -ex

echo "This script automates the steps in the README"

echo "Getting Golang"
go_pkg_remote=https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz
go_pkg_path=./tmp/go-linux-amd64.tar.gz

if [ -f "${go_pkg_path}" ]; then
    echo "Go package already exist, skipping download"
else
    echo "Go package doesn't exist, downloading"
    wget "${go_pkg_remote}" -O "${go_pkg_path}"
fi
echo "${go_pkg_remote}" > ./tmp/go-version.txt

echo "Getting cf"
cli_pkg_remote='https://cli.run.pivotal.io/stable?release=linux64-binary&version=6.31.0&source=github-rel'
cli_pkg_path=./tmp/cf-cli_linux_x86-64.tgz

if [ -f "${cli_pkg_path}" ]; then
    echo "cf cli package already exist, skipping download"
else
    echo "cf cli package doesn't exist, downloading"
    wget "${cli_pkg_remote}" -O "${cli_pkg_path}"
fi
echo "${cli_pkg_remote}" > ./tmp/cli-version.txt

echo "Getting Python"
python_pkg_remote=https://www.python.org/ftp/python/2.7.14/Python-2.7.14.tgz
python_pkg_path=./tmp/python.tgz

if [ -f "${python_pkg_path}" ]; then
    echo "Python package already exist, skipping download"
else
    echo "Python package doesn't exist, downloading"
    wget "${python_pkg_remote}" -O "${python_pkg_path}"
fi
echo "${python_pkg_path}" > ./tmp/python-version.txt

echo "Packaging local source"

tar -cvzf ./tmp/spacebears_src.tgz -C ../src/ spacebears/
tar -cvzf ./tmp/acceptance_tests_src.tgz -C ../src/ acceptance-tests/

echo "Adding blobs"

bosh add-blob "${go_pkg_path}" go-linux-amd64.tar.gz
bosh add-blob ./tmp/go-version.txt go-version.txt

bosh add-blob "${cli_pkg_path}" cf-cli_linux_x86-64.tgz
bosh add-blob ./tmp/cli-version.txt cli-version.txt

bosh add-blob "${python_pkg_path}" python.tgz
bosh add-blob ./tmp/python-version.txt python-version.txt

bosh add-blob ./tmp/spacebears_src.tgz spacebears_src.tgz
bosh add-blob ./tmp/acceptance_tests_src.tgz acceptance_tests_src.tgz

echo "Creating release"
bosh create-release --force

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

yes | bosh -d bosh_simple_with_smoketests deploy manifests/lite_manifest.yml --no-redact
