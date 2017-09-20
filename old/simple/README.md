## simple

Nothing but the app as a BOSH release

## Commands used to bootstrap

```bash
git init

bosh init release simple
cd simple
touch blobs/.gitkeep
git add -A .
git commit -m "Initial release creation"

mkdir tmp
echo "
tmp

config/dev.yml
config/private.yml

.blobs
blobs
.dev_builds
dev_releases
" > .gitignore

echo "Adding go"
export go_verion=1.7.5
wget https://storage.googleapis.com/golang/go$go_verion.linux-amd64.tar.gz -O ./tmp/go-linux-amd64.tar.gz
echo "Go version: $go_verion" > ./tmp/go-version.txt

bosh add blob ./tmp/go-linux-amd64.tar.gz go
bosh add blob ./tmp/go-version.txt go

bosh generate package go
bosh generate package spacebears

mkdir -p src/spacebears
touch src/spacebears/main.go

bosh generate job spacebears-job

echo "Work in in go, spacebears, and src/my-service to build first iteration..." 

bosh create release --force
bosh upload release
bosh deployment manifests/lite.yml

# Iterate...
bosh create release --force && bosh upload release && yes yes | bosh deploy --recreate  --no-redact
```

```bash
nc -w0 -u logs5.papertrailapp.com 25705  <<< "logging from remote"
```
