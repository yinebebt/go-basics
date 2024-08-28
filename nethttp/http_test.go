package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDialer(t *testing.T) {
	t.Run("test dialer with various protocols and services", func(t *testing.T) {
		network := "tcp"

		// Map of protocol/service names to expected ports
		//refer /etc/servcies for more
		protocols := map[string]string{
			"example.com:http":  ":80",
			"example.com:https": ":443",
		}

		for addr, expectedPort := range protocols {
			conn, err := dialer(network, addr)
			assert.Nil(t, err, "Failed to connect to %s", addr)
			if err == nil {
				defer conn.Close()
				assert.Contains(t, conn.RemoteAddr().String(), expectedPort, "Expected port %s for %s", expectedPort, addr)
			}
		}
	})
}
