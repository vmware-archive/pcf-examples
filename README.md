# pcf-examples

This repo builds out several bosh releases to a embedded key-value
database service (spacebears, a toy example based on [bolt](https://github.com/boltdb/bolt)). 
`bosh-simple` is the minimal
[BOSH](https://bosh.io)
release with no dependencies. Each additional
release layers in some other integration or feature. For example, 
`bosh-route-registrar` adds broadcasting a route via Cloud Foundry's Router component.

The `src` directory holds the source code for the database, a 
[service broker](https://github.com/openservicebrokerapi/servicebroker/),
and an example consumer app that uses service created by the broker.

The `tile-*` directories contain the Pivotal Cloud Foundry packaging
for the bosh releases and broker. 

## BOSH Releases

* `bosh-simple`
    - the bare minimum release
* `bosh-route-registrar`
    - `bosh-simple` + route registrar (claims routes in Cloud Foundry)
* `bosh-syslog`
    - `bosh-simple` + syslog logging
* `bosh-metron`
      - `bosh-simple` + metron (emits service metrics)
* `bosh-docker-bosh`
      - spacebears packages as a BOSH deployed docker image
* `bosh-dashboard`
      - `bosh-simple` + sso dashboard for db
* `bosh-odb`
      - adapter for ODB release

## Source code
* `./src/spacebears`
    - Golang key-value store with basic CRUD client API and admin API
* `./src/broker`
    - Python implementation of the Open Service Broker API for Spacebears 
* `./src/sample-app`
    - Python app consumer of Spacebears service.

See [src/README.md](src/README.md) for detailed instructions on how to build,
test, or change source code.

(Spacebears is a nod to CockroachDB.
[Tardigrades](https://en.wikipedia.org/wiki/Tardigrade)
are both more indestructible and cuter and so seemed like a good choice.)

## References

* https://github.com/cloudfoundry/exemplar-release
