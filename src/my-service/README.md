## my-service

* Key value store with a service broker & api

### service-broker
* catalog: single plan
* create-instance
    * create bucket
    * store in some "main" bucket
* bind-instance
	* generate creds
	* store in some "main" bucket
* ubund-instance
    * remove creds from "main" bucket
* delete-instance
    * destroy bucket

### db
* boltdb - filesystem based key value store. multi tenant but no HA
* API (for devs, given out by broker)
    * get to bucket (with auth header)
    * put to bucket (with auth header)
* API for admin (used by SB)
    * create bucket instance
    * create bucket creds...?

### dev
#### Setup

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

#### build
```bash
go generate ./...
go build main.go
```

#### Test
```bash
go generate ./...
go vet ./...
ginkgo -r -failOnPending -race
```

#### Dependency vendoring
via  [govendor](https://github.com/kardianos/govendor)

(Tried `dep`, but it added 10s of megabytes of golang.org/x/... to vendor)

## todo
* come up with clever name
* add makefile to get rid of all the README base code
