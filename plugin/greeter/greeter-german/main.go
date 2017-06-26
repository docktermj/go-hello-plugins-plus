package main

import (
	"github.com/docktermj/go-hello-plugins/interface/greeter"
	"github.com/hashicorp/go-plugin"
)

// ----------------------------------------------------------------------------
// Implementation of interface.
// ----------------------------------------------------------------------------

// Here is a real implementation of Greeter
type GreeterGerman struct{}

func (GreeterGerman) Greet() string {
	return "Grüße aus Deutschland!"
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
	MagicCookieKey:   "greeter-german-cookie-key",
	MagicCookieValue: "greeter-german-cookie-value",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"greeter-german-plugin": &greeter.Plugin{Impl: new(GreeterGerman)},
}

// Run the plugin.
func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
