# go-pmtud

[![GitHub build](https://img.shields.io/github/actions/workflow/status/cunicu/go-pmtud/build.yaml?style=flat-square)](https://github.com/cunicu/go-pmtud/actions)
[![goreportcard](https://goreportcard.com/badge/github.com/cunicu/go-pmtud?style=flat-square)](https://goreportcard.com/report/github.com/cunicu/cugo-pmtudnicu)
[![Codecov](https://img.shields.io/codecov/c/github/cunicu/go-pmtud?token=WWQ6SR16LA&style=flat-square)](https://app.codecov.io/gh/cunicu/go-pmtud)
[![License](https://img.shields.io/github/license/cunicu/go-pmtud?style=flat-square)](https://github.com/cunicu/go-pmtud/blob/master/LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cunicu/go-pmtud?style=flat-square)
[![Go Reference](https://pkg.go.dev/badge/github.com/cunicu/go-pmtud.svg)](https://pkg.go.dev/github.com/cunicu/go-pmtud)

_go-pmtud_ is a Go package which implements portable methods for finding a paths Maximum Transmission Unit (MTU).

## Related RFCs

- **RFC 1191:** [Path MTU Discovery][rfc1191]
- **RFC 8201:** [Path MTU Discovery for IP version 6][rfc8201]
- **RFC 4821:** [Packetization Layer Path MTU Discovery][rfc4821]
- **RFC 8899:** [Packetization Layer Path MTU Discovery for Datagram Transports][rfc8899]
- **RFC 4459:** [MTU and Fragmentation Issues with In-the-Network Tunneling][rfc4459]
- **draft-ietf-tram-stun-pmtud:**: [Packetization Layer Path MTU Discovery (PLPMTUD) For UDP Transports Using Session Traversal Utilities for NAT (STUN)][draft-ietf-tram-stun-pmtud]

[rfc4821]: https://www.rfc-editor.org/rfc/rfc4821.html
[rfc8899]: https://www.rfc-editor.org/rfc/rfc8899
[rfc1191]: https://www.rfc-editor.org/rfc/rfc1191
[rfc8201]: https://www.rfc-editor.org/rfc/rfc8201
[rfc4459]: https://www.rfc-editor.org/rfc/rfc4459
[draft-ietf-tram-stun-pmtud]: https://datatracker.ietf.org/doc/html/draft-ietf-tram-stun-pmtud

## References

- https://wiki.geant.org/display/public/EK/Path+MTU

## Authors

-   Steffen Vogel ([@stv0g](https://github.com/stv0g))

## License

go-pmtud is licensed under the [Apache 2.0](./LICENSE) license.

- SPDX-FileCopyrightText: 2023 Steffen Vogel <post@steffenvogel.de>
- SPDX-License-Identifier: Apache-2.0