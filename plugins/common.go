package plugins

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)


// HtmlCall is the interface that we're exposing as a plugin.
type HtmlCall interface {
	Call(string) string
}

// Greeter is the interface that we're exposing as a plugin.
type Greeter interface {
	Greet() string
}

// Here is an implementation that talks over RPC
type HtmlCallRPC struct{ client *rpc.Client }

// Here is an implementation that talks over RPC
type GreeterRPC struct{ client *rpc.Client }


func (g *HtmlCallRPC) Call(data string) string {
	var resp string
	err := g.client.Call("Plugin.Call", data, &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}
	return resp
}

func (g *GreeterRPC) Greet() string {
	var resp string
	err := g.client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}


type HtmlCallRPCServer struct {
	Impl HtmlCall
}

func (s *HtmlCallRPCServer) Call(data string, resp *string) error {
	*resp = s.Impl.Call(data)
	return nil
}

// Here is the RPC server that GreeterRPC talks to, conforming to
// the requirements of net/rpc
type GreeterRPCServer struct {
	// This is the real implementation
	Impl Greeter
}

func (s *GreeterRPCServer) Greet(args interface{}, resp *string) error {
	*resp = s.Impl.Greet()
	return nil
}

// This is the implementation of plugin.Plugin so we can serve/consume this
//
// This has two methods: Server must return an RPC server for this plugin
// type. We construct a GreeterRPCServer for this.
//
// Client must return an implementation of our interface that communicates
// over an RPC client. We return GreeterRPC for this.
//
// Ignore MuxBroker. That is used to create more multiplexed streams on our
// plugin connection and is a more advanced use case.
type GreeterPlugin struct {
	// Impl Injection
	Impl Greeter
}

func (p *GreeterPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &GreeterRPCServer{Impl: p.Impl}, nil
}

func (GreeterPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GreeterRPC{client: c}, nil
}


// ****************  MY IMPL

type HtmlCallPlugin struct {
	// Impl Injection
	Impl HtmlCall
}

func (p *HtmlCallPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &HtmlCallRPCServer{Impl: p.Impl}, nil
}

func (HtmlCallPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &HtmlCallRPC{client: c}, nil
}
