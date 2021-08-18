package plugin

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type Plugin interface {
	// Init initialize plugin with specified configuration.
	// If error occurs in initialization, don't return, just PANIC!
	Init(conf string)

	// Serve http request
	ServeHTTP(ctx *fasthttp.RequestCtx)
}

var plugins = make(map[string]Plugin)
var pluginsReady = make(map[string]Plugin)

// Register makes a plugin available by the provided name.
// Plugins are not usable until they have been intialized.
func Register(name string, plugin Plugin) {
	if plugin == nil {
		panic("plugin is nil")
	}
	if _, ok := plugins[name]; ok {
		panic("register called twice for plugin " + name)
	}
	plugins[name] = plugin
}

// Open init a named plugin with specified configuration.
// After opened, a plugin is ready to serve http request.
func Open(name, conf string) {
	var ok bool
	var plugin Plugin
	if plugin, ok = pluginsReady[name]; ok {
		panic("open called twice for plugin " + name)
	}
	if plugin, ok = plugins[name]; !ok {
		panic("no register has been called for plugin " + name)
	}
	plugin.Init(conf)
	pluginsReady[name] = plugin
}

// Dispatch uri location to a plugin with specified name:
// location => name => plugin
func Dispatch(location, name string, router *router.Router) {
	var ok bool
	var plugin Plugin
	if plugin, ok = pluginsReady[name]; !ok {
		panic("no available plugin for " + name)
	}

	router.GET(location, plugin.ServeHTTP)
	router.POST(location, plugin.ServeHTTP)
}
