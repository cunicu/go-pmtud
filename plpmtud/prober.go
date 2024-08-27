// SPDX-FileCopyrightText: 2023-2024 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package plpmtud

type Prober interface {
	SendProbeRequest(mtu uint) error
	SendProbeResponse(mtu uint) error

	RegisterDiscoverer(h *Discoverer)

	Close() error
}
