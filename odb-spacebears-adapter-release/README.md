
bosh add-blob target/odb-spacebears-adapter-release-0.0.1-SNAPSHOT.jar adapter.jar
bosh create-release --name=odb-spacebears-adapter --force
bosh upload-release --name=odb-spacebears-adapter
bosh -d odb-spacebears deploy manifests/odb.yml --no-redact