# pcf-examples

This repo builds out several bosh releases to a embedded key-value
database service (spacebears, a toy example based on [bolt](https://github.com/boltdb/bolt)). `bosh-simple` is the minimal [BOSH](https://bosh.io)
release with no dependencies. Each additional release layers in some other integration or feature. For example, `bosh-route-registrar` adds broadcasting a route via Cloud Foundry's Router component.

The `src` directory holds the source code for the database, a [service broker](https://github.com/openservicebrokerapi/servicebroker/), and an example consumer app that uses service created by the broker.

The `tile-*` directories contain the Pivotal Cloud Foundry packaging
for the bosh releases and broker.

## Additional Resources

* [BOSH Addon Development Plan](addon-development-plan.md)
* [Troubleshooting Guide](troubleshooting-guide.md)

## BOSH Releases

* [bosh-simple](bosh-simple-addon/README.md)
  * the bare minimum release

* [bosh-simple-with-link](bosh-simple-with-link/README.md)
  * `bosh-simple` + BOSH link for sharing address and credentials

* [bosh-simple-with-routing](bosh-simple-bosh-simple-with-routing/README.md)
  * `bosh-simple` + route registrar (claims routes in Cloud Foundry)
* [bosh-simple-with-smoketests](bosh-simple-bosh-simple-bosh-simple-with-smoketests/README.md)
  * `bosh-simple` + smoketests

* [odb-spacebears-adapter-release](odb-spacebears-adapter-release/README.md)
  * Simple On Demand Broker BOSH release of Spacebears with a Java service adapter.

* [java-service-broker](java-service-broker/readme.md)
  * Simple Java implementation of the Open Service Broker API

## Tiles

* [tile-for-bosh-simple](tile-for-bosh-simple/README.md)
* [tile-for-bosh-simple-addon](tile-for-bosh-simple-addon/README.md)
* [tile-for-bosh-simple-with-link](tile-for-bosh-simple-with-link/README.md)
* [tile-for-bosh-simple-with-routing](tile-for-bosh-simple-with-routing/README.md)

## Source code

* `src/broker`
  * Python implementation of the Open Service Broker API for Spacebears

* `src/db-docker`
  * packages spacebears as a docker image

* `src/db-helm`
  * spacebears service and spacebears broker helm chart

* `src/sample-app`
  * Python app consumer of Spacebears service.

* `src/spacebears`
  * Golang key-value store with basic CRUD client API and admin API

See [src/README.md](src/README.md) for detailed instructions on how to build,
test, or change source code.

(Spacebears is a nod to CockroachDB.
[Tardigrades](https://en.wikipedia.org/wiki/Tardigrade)
are both more indestructible and cuter and so seemed like a good choice.)

## References

* [Cloud Foundry Exemplar Release](https://github.com/cloudfoundry/exemplar-release)
