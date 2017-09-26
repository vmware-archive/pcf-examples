## spacebears
Filesystem based key-value store. This is multi-tenant, but not highly available. Clustering is not possible. 
Just an example, don't use it for anything remotely real.
* Client API for developers
    - CRUD operations on keys within a bucket
* Admin API (used by the service broker) to 
    - create and delete the buckets.
    - generate credentails for the client API.

### Setup

The Go code doesn't use full url qualified package paths. Add the project root
to the `GOPATH` with

```bash
export GOPATH=$GOPATH:$(pwd)
```

Install dependencies (test and vendor)
```bash
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
go get github.com/maxbrunsfeld/counterfeiter
go get github.com/kardianos/govendor
```

### Build
```bash
go generate ./...
go build main.go
```

### Unit Test
```bash
go generate ./...
go vet ./...
ginkgo -r -failOnPending -race
```

### Run
```bash
export ADMIN_USERNAME=admin
export ADMIN_PASSWORD=password
export PORT=9000

go run main.go
```

#### Dependency vendoring
To change dependencies, see [govendor](https://github.com/kardianos/govendor) docs for specific commands.

(Tried `dep`, but it added 10s of megabytes of golang.org/x/... to vendor)

## service-broker
Implementation of [service broker API](https://github.com/openservicebrokerapi/servicebroker/) to Spacebears.

### Setup
The application is written for Python 3.

Install dependencies
```bash
pip install -r requirements.txt
```

### Unit Test
```bash
python -m unittest discover -v -s . -p '*_test.py'
```

### Run
```bash
export ADMIN_USERNAME=admin
export ADMIN_PASSWORD=awesome_broker
export DB_ADMIN_USERNAME=admin
export DB_ADMIN_PASSWORD=password
export DB_URL=http://localhost:9000

python broker.py
```
## Integration testing (via PCF Dev)
Testing all these things together requires:
* `spacebears` is running (ideally as a BOSH release)
* `service-broker` is running (ideally pushed as a Cloud Foundry app)
* `sample-app` is running on Pivotal Cloud Foundry

#### Run spacebears
To run Spacebears as a BOSH release locally, see `bosh-simple`.

#### Run service-broker
To run the service broker on a local Cloud Foundry, see the
[cf-deployment](https://github.com/cloudfoundry/cf-deployment)
run instructions. Deploy this to the same bosh-lite as the `bosh-simple` deployment.

Target the cf-deployment's cf
(get the cf-deployment password with `grep "cf_admin_password" deployment-vars.yml`)
```bash
cf login -a https://api.bosh-lite.com --skip-ssl-validation -u admin -p XXXXXXXXX
```

Allow all apps ton talk to the Spacebears db node (only needed if talking via IP and haven't registered a route)
```bash
cd src/broker
cf create-security-group spacebears-asg spacebears-asg.json
cf bind-running-security-group spacebears-asg
```

Then push the broker as an app and add it to the marketplace
```bash
cd src/broker
cf push
cf create-service-broker spacebears-broker admin monkey123 https://spacebears-broker.bosh-lite.com
cf enable-service-access spacebears-db
```

View the logs with
```bash
echo "Get recent logs"
cf logs spacebears-broker --recent
echo "Stream new logs"
cf logs spacebears-broker
```

#### Run sample-app
To run the sample app

```bash
cd src/sample-app
cf create-service spacebears-db plan1 my-spacebears
cf push
```

View the logs with
```bash
echo "Get recent logs"
cf logs spacebears-consumer --recent
echo "Stream new logs"
cf logs spacebears-consumer
```

Visit http://spacebears-consumer.local.pcfdev.io/ to verify success
