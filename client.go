package tpshp

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
)

// Client implements a client for the TP-Link Smart-Home protocol
type Client interface {
	// Call sends one or more commands to the TP-Link device
	Call(context.Context, *Request) error
}

// New creates a new TP-Link Smart-Home client for the given IP address. The port defaults to 9999.
// If you need to use a different port, see NewWithPort()
func New(ip string) Client {
	return NewWithPort(ip, 9999)
}

// NewWithPort creates a new TP-Link Smart-Home client for the given IP address and port
func NewWithPort(ip string, port uint16) Client {
	return &client{
		ip:   ip,
		port: port,
	}
}

// client implements the Client interface
type client struct {
	ip   string
	port uint16
}

// call sens one or more command requests to the TP-Link device
func (cli *client) Call(ctx context.Context, req *Request) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return err
	}

	// connect to the device
	conn, err := (&net.Dialer{}).DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", cli.ip, cli.port))
	if err != nil {
		return err
	}

	if err := SendRaw(conn, payload); err != nil {
		return err
	}

	if req.ResponseExpected() {
		response, err := RecvRaw(conn)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(response, req); err != nil {
			return err
		}
	}

	return nil
}
