package plugins

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"net/http"
)



func PluginHandler(response http.ResponseWriter, request *http.Request)() {
	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("sample-plugins/console-echo/console-echo.exe"),
		Logger:          logger,
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("greeter")
	if err != nil {
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	greeter := raw.(Greeter)
	fmt.Println(greeter.Greet())


	// Request a second plugin
	raw, err = rpcClient.Dispense("marryme")
	if err != nil {
		log.Fatal(err)
	}
	greeter = raw.(Greeter)
	fmt.Println(greeter.Greet())


	raw, err = rpcClient.Dispense("call")
	if err != nil {
		log.Fatal(err)
	}
	call := raw.(HtmlCall)
	response.Write([] byte(call.Call("F")))

}


func CleanPluginHandler(response http.ResponseWriter, request *http.Request)() {

	// pluginMap is the map of plugins we can dispense.
	var cleanPluginMap = map[string]plugin.Plugin{
		"call": &HtmlCallPlugin{},
		"doublecall": &HtmlCallPlugin{},
	}

	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: GetHandshakeConfig(),
		Plugins:         cleanPluginMap,
		Cmd:             exec.Command("sample-plugins/simple-plugin/simple-plugin.exe"),
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	raw, err := rpcClient.Dispense("doublecall")
	if err != nil {
		log.Fatal(err)
	}
	call := raw.(HtmlCall)
	response.Write([] byte(call.Call("F")))

}

func GetHandshakeConfig() plugin.HandshakeConfig {
	return plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "BASIC_PLUGIN",
		MagicCookieValue: "hello",
	}
}

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"greeter": &GreeterPlugin{},
	"marryme": &GreeterPlugin{},
	"call": &HtmlCallPlugin{},
}