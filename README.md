mm-ctl
======

MidoNet port binding command line tool

Prerequisites
-------------

* [Go](https://golang.org/dl/) (>= 1.5)

mm-ctl depends on the specific versions of the libraries. Especially mm-ctl
heavily depends on [curator.go][curator]'s [`InterProcessMutex`][ipm] to have
the coherence with the lock of MidoNet agent against the virutal topology, and
curator.go is not up-to-date to be compatible with the latest version of
[go-zookeeper][zk]. So the snapshots of dependencies are included in the
codebase under `vendor` directory. Go 1.5 introduced [the vendoring][vendoring]
as its experimental feature and we leverage it for managing the dependencies.

[curator]: https://github.com/flier/curator.go
[ipm]: https://godoc.org/github.com/flier/curator.go/recipes#InterProcessMutex
[zk]: https://github.com/samuel/go-zookeeper
[vendoring]: https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit

Installation
------------

The following commands generate the binary `mm-ctl`. You can use it directly or
put anywhere you want.

```
$ GO15VENDOREXPERIMENT=1 go get
$ GO15VENDOREXPERIMENT=1 go build
```

Usage
-----

mm-ctl assumes two config files are placed in the appropriate places,
`/etc/midolman/midolman.conf` and `/etc/midonet_host_id.properties`. If you
have them in the different paths, you can specify them with `-config` and
`-host-config` for each other.

### Binding the Neutron port to the interface

```
$ sudo mm-ctl bind 2738e706-6ec7-4b22-b43f-be20f7ad85cb b3f9a34d-veth
```

or

```
$ sudo mm-ctl -config /path/to/midolman.conf \
  -host-config /path/to/midonet_host_id.properties \
  bind 2738e706-6ec7-4b22-b43f-be20f7ad85cb b3f9a34d-veth
```

### Unbinding the Neutron port

```
$ sudo mm-ctl unbind 2738e706-6ec7-4b22-b43f-be20f7ad85cb
```

or

```
$ sudo mm-ctl -config /path/to/midolman.conf \
  -host-config /path/to/midonet_host_id.properties \
  unbind 2738e706-6ec7-4b22-b43f-be20f7ad85cb
```


License
-------

mm-ctl is released under Apache License, Version 2.0. See `LICENSE` for more
details.
