package main

import (
	"github.com/docktermj/go-hello-plugins/interface/greeter"
	"github.com/hashicorp/go-plugin"
)

// ----------------------------------------------------------------------------
// Implementation of interface.
// ----------------------------------------------------------------------------

type GreeterItalian struct{}

func (GreeterItalian) Greet() string {
	return "saluti DALL'ITALIA"
}

// ----------------------------------------------------------------------------
// Install and run plugin.
// ----------------------------------------------------------------------------

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "greeter-italian-cookie-key",
	MagicCookieValue: "greeter-italian-cookie-value",
}

var pluginMap = map[string]plugin.Plugin{
	"greeter-italian-plugin": &greeter.Plugin{Impl: new(GreeterItalian)},
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
