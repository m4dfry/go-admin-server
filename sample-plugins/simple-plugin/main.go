package main

import (
	"github.com/hashicorp/go-plugin"
	common "github.com/m4dfry/go-admin-server/plugins"
)

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"call": &common.HtmlCallPlugin{},
	"doublecall": &common.HtmlCallPlugin{},
}

type  AnswerPlugin struct {}

func (g* AnswerPlugin) Call(str string) string {
	return "{" + str + ":call_object}"
}

type  DoubleCall struct {}

func (g* DoubleCall) Call(str string) string {
	return "{" + str + ":double_call_object}"
}

func main() {
	htmlcall := &AnswerPlugin{}
	doublecall := &DoubleCall{}
	pluginMap["call"] = &common.HtmlCallPlugin{Impl: htmlcall}
	pluginMap["doublecall"] = &common.HtmlCallPlugin{Impl: doublecall}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: common.GetHandshakeConfig(),
		Plugins:         pluginMap,
	})
}

