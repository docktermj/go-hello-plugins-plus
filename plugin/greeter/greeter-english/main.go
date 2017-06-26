package main

import (
	"github.com/docktermj/go-hello-plugins/interface/greeter"
	"github.com/hashicorp/go-plugin"
)

// ----------------------------------------------------------------------------
// Implementation of interface.
// ----------------------------------------------------------------------------

type GreeterEnglish struct{}

func (GreeterEnglish) Greet() string {
	return "Greetings from U.S.A.!"
}

// ----------------------------------------------------------------------------
// Install and run plugin.
// ----------------------------------------------------------------------------

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "greeter-english-cookie-key",
	MagicCookieValue: "greeter-english-cookie-value",
}

var pluginMap = map[string]plugin.Plugin{
	"greeter-english-plugin": &greeter.Plugin{Impl: new(GreeterEnglish)},
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
