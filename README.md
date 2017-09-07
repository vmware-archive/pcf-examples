## Bosh Examples

A few small
[BOSH](https://bosh.io)
examples.

See src/README.md for service & broker code details

## todo

* Base service: key/value db + service broker
    * Add make file to deal with Go commands
* BOSH release only
* tile-generator BOSH release
* BOSH release + route registrar
* BOSH release + syslog
* Docker-ify service + docker-bosh
* BOSH release + metron
* BOSH release + dashboard / sso
* ODB-ify BOSH release?

Tree
* `./old` Delete this when done copying useful coe. Some examples from before, half working.
* `./src` Source dir. Go code under here, using as root to import path.
* `./src/my-service` Go implementation of key-value store, service
* `./src/broker` Python service broker

## References

* https://github.com/cloudfoundry/exemplar-release
