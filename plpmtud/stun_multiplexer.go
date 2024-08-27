// SPDX-FileCopyrightText: 2023 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package plpmtud

import (
	"fmt"
	"net"

	"github.com/pion/stun/v3"
)

const (
	StunMethodFallthrough stun.Method = 0xffff
)

type StunHandler func(*stun.Message) error

type StunMultiplexer struct {
	*net.UDPConn

	methodHandlers map[stun.Method]StunHandler
}

func NewStunMultiplexer(c *net.UDPConn) (*StunMultiplexer, error) {
	sc := &StunMultiplexer{
		UDPConn: c,

		methodHandlers: map[stun.Method]StunHandler{},
	}

	return sc, nil
}

func (c *StunMultiplexer) RegisterStunHandler(m stun.Method, h StunHandler) {
	c.methodHandlers[m] = h
}

func (c *StunMultiplexer) ReadFromUDP(b []byte) (int, *net.UDPAddr, error) {
	for {
		n, addr, err := c.UDPConn.ReadFromUDP(b)
		if err != nil {
			return -1, nil, err
		}

		if stun.IsMessage(b) {
			msg := &stun.Message{
				Raw: b,
			}

			if err := msg.Decode(); err != nil {
				return 0, nil, fmt.Errorf("failed to decode message: %w", err)
			}

			if h, ok := c.methodHandlers[msg.Type.Method]; ok {
				err = h(msg)
			} else if h, ok := c.methodHandlers[StunMethodFallthrough]; ok {
				err = h(msg)
			}
			if err != nil {
				return 0, nil, fmt.Errorf("failed to invoke handlers: %w", err)
			}
		} else {
			return n, addr, nil
		}
	}
}

func (c *StunMultiplexer) WriteStunMessage(msg *stun.Message) error {
	msg.Encode()

	_, err := c.Write(msg.Raw)

	return err
}
