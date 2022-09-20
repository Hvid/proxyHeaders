package proxyHeaders

import (
	"context"
	"net"
	"net/http"
)

type Config struct {
	Enabled bool
}

func CreateConfig() *Config {
	return &Config{
		Enabled: true,
	}
}

type Plugin struct {
	enabled bool
	name    string
	next    http.Handler
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Plugin{
		enabled: config.Enabled,
		name:    name,
		next:    next,
	}, nil
}

func (plugin *Plugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if !plugin.enabled {
		plugin.next.ServeHTTP(rw, req)
		return
	}
	var ip string
	var port string
	ip, port = silentSplitHostPort(req.RemoteAddr)
	req.Header.Set("TRAEFIK-SRCPORT", port)
	req.Header.Set("TRAEFIK-SRCIP", ip)
	plugin.next.ServeHTTP(rw, req)
}

func silentSplitHostPort(value string) (host, port string) {
	host, port, err := net.SplitHostPort(value)
	if err != nil {
		return value, "-"
	}
	return host, port
}
