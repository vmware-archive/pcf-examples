## bosh-simple-with-link

This adds a [BOSH link](https://bosh.io/docs/links.html) to share the
address and credentials for the Spacebears service with other BOSH
jobs in a deployment. This allows the service to be used inside a
private network, without exposing a route (cf.
[bosh-simple-with-routing](../bosh-simple-with-routing).

The only change to add this BOSH link is the additional `provides`
section in the [spacebears_job spec file](jobs/spacebears_job/spec).

For more information about BOSH links, see
[BOSH Links: Why and How](https://gist.github.com/Amit-PivotalLabs/c39528248b8cdc4ba8e347f8aa68abb6)
and the
[official BOSH links documentation](https://bosh.io/docs/links.html).

### Setting up release

(`deploy.sh` automates these steps)

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
tar -cvzf ./tmp/spacebears_src.tgz -C ../src/ spacebears/

bosh add-blob ./tmp/spacebears_src.tgz spacebears_src.tgz
```

Create & upload release
```bash
bosh create-release --force
bosh upload-release
```

### Deploy (Lite)

If this is a fresh environment, be sure to upload a proper stemcell
```bash
bosh upload-stemcell https://s3.amazonaws.com/bosh-core-stemcells/warden/bosh-stemcell-3445.7-warden-boshlite-ubuntu-trusty-go_agent.tgz
```

The bosh deployment manifest in `manifest/lite_manifest.yml` is setup to work
with a default [deployed BOSH Lite](https://bosh.io/docs/bosh-lite) using
the [bosh-deployment warden cloud-config](https://github.com/cloudfoundry/bosh-deployment/blob/master/warden/cloud-config.yml)

Deploy
```bash
bosh -d bosh_simple deploy manifests/lite_manifest.yml --no-redact
```

#### Troubleshooting
```bash
bosh -d bosh_simple ssh spacebears_db_node
# hope over to root for monit and other commands
sudo su -
```

* job logs
    * `/var/vcap/sys/log/spacebears/spacebears.out.log`
    * `/var/vcap/sys/log/spacebears/spacebears.err.log`
* monit logs
    * `/var/vcap/monit/monit.log`

###  Cleanup

```bash
bosh -d bosh_simple delete-deployment --force
bosh delete-release bosh-simple-spacebears
bosh clean-up
```
