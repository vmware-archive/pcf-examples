bosh create-release --force --name dummy-adapter
bosh upload-release --name dummy-adapter
bosh -d spacebears-on-demand-broker-dev deploy spacebears_manifest.yml
