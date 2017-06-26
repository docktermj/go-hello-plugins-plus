package main

import (
	"github.com/docktermj/go-hello-plugins/interface/hello"
	"github.com/hashicorp/go-plugin"
)

// ----------------------------------------------------------------------------
// Implementation of interface.
// ----------------------------------------------------------------------------

type HelloEnglish struct{}

func (HelloEnglish) Speak() string {
	return "Hello World!"
}

// ----------------------------------------------------------------------------
// Install and run plugin.
// ----------------------------------------------------------------------------

// Information to verify correct plugin.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "hello-english-cookie-key",
	MagicCookieValue: "hello-english-cookie-value",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"hello-english-plugin": &hello.Plugin{Impl: new(HelloEnglish)},
}

// Run the plugin.
func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
