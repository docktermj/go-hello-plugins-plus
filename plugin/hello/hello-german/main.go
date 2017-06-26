package main

import (
	"github.com/docktermj/go-hello-plugins/interface/hello"
	"github.com/hashicorp/go-plugin"
)

// ----------------------------------------------------------------------------
// Implementation of interface.
// ----------------------------------------------------------------------------

type HelloGerman struct{}

func (HelloGerman) Speak() string {
	return "Hallo, Welt!"
}

// ----------------------------------------------------------------------------
// Install and run plugin.
// ----------------------------------------------------------------------------

// Information to verify correct plugin.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "hello-german-cookie-key",
	MagicCookieValue: "hello-german-cookie-value",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"hello-german-plugin": &hello.Plugin{Impl: new(HelloGerman)},
}

// Run the plugin.
func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
