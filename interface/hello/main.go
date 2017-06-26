package hello

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// ----------------------------------------------------------------------------
// Interface
// ----------------------------------------------------------------------------

type Hello interface {
	Speak() string
}

// ----------------------------------------------------------------------------
// RPC
// ----------------------------------------------------------------------------

type RPC struct{ client *rpc.Client }

func (g *RPC) Speak() string {
	var resp string
	err := g.client.Call("Plugin.Speak", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}
	return resp
}

// ----------------------------------------------------------------------------
// RPCServer
// ----------------------------------------------------------------------------

type RPCServer struct {
	Impl Hello
}

func (s *RPCServer) Speak(args interface{}, resp *string) error {
	*resp = s.Impl.Speak()
	return nil
}

// ----------------------------------------------------------------------------
// Plugin
// ----------------------------------------------------------------------------

type Plugin struct {
	Impl Hello
}

func (p *Plugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: p.Impl}, nil
}

func (Plugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPC{client: c}, nil
}
