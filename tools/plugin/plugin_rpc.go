package plugin

import (
	"log"

	"github.com/smallnest/rpcx/server"
)

type PluginRPC interface {
	// Init initialize plugin with specified configuration.
	// If error occurs in initialization, don't return, just PANIC!
	Init(conf string)

	/*
		Serve rpc可注册的方法:
		必须是可导出类型的方法
		接受3个参数，第一个是 context.Context类型，其他2个都是可导出（或内置）的类型。
		第3个参数是一个指针
		有一个 error 类型的返回值
		建议定义方法:
		ServeRPC(ctx context.Context, args *pb.Request, reply *pb.Response) error
	*/

}

var plugins_RPC = make(map[string]PluginRPC)
var pluginsReady_RPC = make(map[string]PluginRPC)

// Register makes a plugin available by the provided name.
// Plugins are not usable until they have been intialized.
func RegisterRPC(name string, plugin PluginRPC) {
	if plugin == nil {
		panic("plugin is nil")
	}
	if _, ok := plugins_RPC[name]; ok {
		panic("register called twice for plugin " + name)
	}
	plugins_RPC[name] = plugin
}

// Open init a named plugin with specified configuration.
// After opened, a plugin is ready to serve http request.
func OpenRPC(name, conf string) {
	var ok bool
	var plugin PluginRPC
	if plugin, ok = pluginsReady_RPC[name]; ok {
		panic("open called twice for plugin " + name)
	}
	if plugin, ok = plugins_RPC[name]; !ok {
		panic("no register has been called for plugin " + name)
	}
	plugin.Init(conf)
	pluginsReady_RPC[name] = plugin
}

// Dispatch specified name to a plugin:
// name => plugin rpc
func DispatchRPC(name string, rpc *server.Server) {
	var ok bool
	var plugin PluginRPC
	if plugin, ok = pluginsReady_RPC[name]; !ok {
		panic("no available plugin for " + name)
	}
	if err := rpc.RegisterName(name, plugin, ""); err != nil {
		log.Fatal(err)
	}
}
