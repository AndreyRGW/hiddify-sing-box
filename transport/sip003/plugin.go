package sip003

import (
	"context"
	"net"
	"sync"

	"github.com/sagernet/sing-box/adapter"
	E "github.com/sagernet/sing/common/exceptions"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
)

type PluginConstructor func(ctx context.Context, pluginArgs Args, router adapter.Router, dialer N.Dialer, serverAddr M.Socksaddr) (Plugin, error)

type Plugin interface {
	DialContext(ctx context.Context) (net.Conn, error)
}

var (
	plugins   map[string]PluginConstructor
	pluginsMu sync.Mutex
)

func RegisterPlugin(name string, constructor PluginConstructor) {
	pluginsMu.Lock()
	defer pluginsMu.Unlock()
	if plugins == nil {
		plugins = make(map[string]PluginConstructor)
	}
	plugins[name] = constructor
}

func CreatePlugin(ctx context.Context, name string, pluginArgs string, router adapter.Router, dialer N.Dialer, serverAddr M.Socksaddr) (Plugin, error) {
	pluginOptions, err := ParsePluginOptions(pluginArgs)
	if err != nil {
		return nil, E.Cause(err, "parse plugin_opts")
	}
	pluginsMu.Lock()
	constructor, loaded := plugins[name]
	pluginsMu.Unlock()
	if !loaded {
		return nil, E.New("plugin not found: ", name)
	}
	return constructor(ctx, pluginOptions, router, dialer, serverAddr)
}
