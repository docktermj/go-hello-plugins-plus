package main

import (
	"github.com/docktermj/go-hello-plugins/interface/hello"
	"github.com/hashicorp/go-plugin"
)

// ----------------------------------------------------------------------------
// Implementation of interface.
// ----------------------------------------------------------------------------

// Here is a real implementation of Greeter
type HelloItalian struct{}

func (HelloItalian) Speak() string {
	return "Ciao mondo!"
}

// ----------------------------------------------------------------------------
// Install and run plugin.
// ----------------------------------------------------------------------------

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "hello-italian-cookie-key",
	MagicCookieValue: "hello-italian-cookie-value",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"hello-italian-plugin": &hello.Plugin{Impl: new(HelloItalian)},
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
