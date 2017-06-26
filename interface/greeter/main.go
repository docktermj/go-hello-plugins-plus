package greeter

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// ----------------------------------------------------------------------------
// Interface
// ----------------------------------------------------------------------------

// Greeter is the interface that we're exposing as a plugin.
type Greeter interface {
	Greet() string
}

// ----------------------------------------------------------------------------
// RPC
// ----------------------------------------------------------------------------

// Here is an implementation that talks over RPC
type RPC struct{ client *rpc.Client }

func (g *RPC) Greet() string {
	var resp string
	err := g.client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

// ----------------------------------------------------------------------------
// RPCServer
// ----------------------------------------------------------------------------

// Here is the RPC server that RPC talks to, conforming to
// the requirements of net/rpc
type RPCServer struct {
	// This is the real implementation
	Impl Greeter
}

func (s *RPCServer) Greet(args interface{}, resp *string) error {
	*resp = s.Impl.Greet()
	return nil
}

// ----------------------------------------------------------------------------
// Plugin
// ----------------------------------------------------------------------------

// This is the implementation of plugin.Plugin so we can serve/consume this
//
// This has two methods: Server must return an RPC server for this plugin
// type. We construct a RPCServer for this.
//
// Client must return an implementation of our interface that communicates
// over an RPC client. We return RPC for this.
//
// Ignore MuxBroker. That is used to create more multiplexed streams on our
// plugin connection and is a more advanced use case.
type Plugin struct {
	// Impl Injection
	Impl Greeter
}

func (p *Plugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: p.Impl}, nil
}

func (Plugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPC{client: c}, nil
}
