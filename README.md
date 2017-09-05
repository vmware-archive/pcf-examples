## Bosh Examples

A few small
[BOSH](https://bosh.io)
examples.

See src/README.md for service & broker code details

## todo

* Base service: key/value db + service broker
    * Add make file to deal with Go commands
* Just the BOSH release
* tile-generator BOSH release
* BOSH release + metron
* BOSH release + syslog
* BOSH release + route registrar
* BOSH release + dashboard / sso
* Docker-ify service + docker-bosh

Tree
* `./old` Delete this when done copying useful coe. Some examples from before, half working.
* `./src` Source dir. Go code under here, using as root to import path.
* `./src/my-service` Go implementation of key-value store, service
* `./src/broker` Python service broker
