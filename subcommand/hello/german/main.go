package german

// Reference: https://github.com/hashicorp/go-plugin/blob/master/examples/basic/main.go

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/docktermj/go-hello-plugins/interface/hello"
	"github.com/hashicorp/go-plugin"
)

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "hello-german-cookie-key",
	MagicCookieValue: "hello-german-cookie-value",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"hello-german-plugin": &hello.Plugin{},
}

// The command.
func Command(argv []string) {

	// We don't want to see the plugin logs.
	log.SetOutput(ioutil.Discard)

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("hello-german"),
	})
	defer client.Kill()

	// Connect via RPC.
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin.
	raw, err := rpcClient.Dispense("hello-german-plugin")
	if err != nil {
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	greeter := raw.(hello.Hello)
	fmt.Println(greeter.Speak())
}
