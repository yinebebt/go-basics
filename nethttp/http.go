// Package nethttp demonstrates how to create an HTTP server in Go.
package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {
	dialer := &net.Dialer{
		Timeout: 5 * time.Second,
	}

	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			fmt.Printf("network: %v addr: %v", network, addr)
			conn, err := dialer.DialContext(ctx, network, addr)
			if err != nil {
				return nil, err
			}
			fmt.Printf("Local Address: %v\n", conn.LocalAddr())   // local ip address of machine assigned by the router
			fmt.Printf("Remote Address: %v\n", conn.RemoteAddr()) // ip address of the service/server we are dialing, in this case the addr's ip address
			return conn, nil
		},
	}

	client := &http.Client{
		Transport: transport,
	}

	_, err := client.Get("https://example.com")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func dialer(ntw, addr string) (net.Conn, error) {
	dialer, err := net.Dial(ntw, addr)
	if err != nil {
		err = fmt.Errorf("unable to dial on address: %v\n, err:%v", ":http", err)
		return nil, err
	}
	dialer.RemoteAddr()
	fmt.Printf("succesafully dialed, \n\tlocal address: %v \n\tremote address: %v\n", dialer.LocalAddr(), dialer.RemoteAddr())
	return dialer, nil
}
