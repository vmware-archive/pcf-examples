
The syslog-release has the actual BOSH release code that we're after. In the context of tiles, though
we'll need the release to be packaged into a tarball so it can be added to the tile.
 
```bash
echo "Clone the syslog release locally"
git clone https://github.com/cloudfoundry/syslog-release ./tmp/syslog-release

echo "Create a tarball from the most recent release"
./tmp/syslog-release
bosh create release releases/syslog/syslog-10.yml --with-tarball

echo "Upload the release to the bosh director"
bosh upload release ./releases/syslog/syslog-10.tgz 

echo "Configure the manifest to include the new properties needed for syslog"

vim manifests/lite.yml

echo "Set deployment"
bosh deployment manifests/lite.yml

# Iterate...
bosh create release --force && bosh upload release && yes yes | bosh deploy --no-redact --recreate
```

# todo

Validate this approach in a tile, ie, script packaging up 2 bosh releases into a manifest, etc