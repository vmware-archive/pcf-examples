
### Setting up release

Get the Golang distribution and add to blobs

```bash
export go_pkg_remote=https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz

wget "${go_pkg_remote}" -O ./tmp/go-linux-amd64.tar.gz
echo "${go_pkg_remote}" > ./tmp/go-version.txt

bosh add-blob ./tmp/go-linux-amd64.tar.gz go-linux-amd64.tar.gz
bosh add-blob ./tmp/go-version.txt go-version.txt   
```

Package up our source code and add to blobs

```bash
tar -cvzf ./tmp/spacebears_src.tgz -C ../src/ my-service/

bosh add-blob ./tmp/spacebears_src.tgz spacebears_src.tgz
```

Create & upload release

```bash
bosh create-release --force
bosh upload-release
```

### Deploy

The bosh deployment manifest in `manifest/gcp_manifest.yml` is setup to work
with a [bosh-bootloader](https://github.com/cloudfoundry/bosh-bootloader) deployed
BOSH on GCP (using the default cloud-config)

Deploy

```bash
bosh upload-stemcell https://s3.amazonaws.com/bosh-gce-light-stemcells/light-bosh-stemcell-3445.7-google-kvm-ubuntu-trusty-go_agent.tgz
bosh -d bosh_only deploy manifests/gcp_manifest.yml --no-redact
```
