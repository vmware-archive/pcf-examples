# tile-for-bosh-simple-with-link

This directory contains the resources to build
the broker and Spacebears' BOSH release into a
[tile](https://docs.pivotal.io/tiledev/tile-structure.html).
The broker obtains the Spacebears service's IP address and credentials
via a BOSH link. This is a good pattern to follow if you do not want
your service exposed via a route (cf.
[tile-for-bosh-simple-with-routing](../tile-for-bosh-simple-with-routing).

The tile is built via
[tile-generator](https://github.com/cf-platform-eng/tile-generator/).

`package.sh` automates building the tile.
