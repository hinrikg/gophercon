package webserver

import "net"
import "net/http"

type WebServer struct {
	http.Server
}

func New(host, port string, handler http.Handler) *WebServer {
	var server WebServer

	server.Addr = net.JoinHostPort(host, port)
	server.Handler = handler

	return &server
}

func (server *WebServer) Start() error {
	return server.ListenAndServe()
}
