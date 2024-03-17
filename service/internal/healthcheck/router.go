package healthcheck

import (
	"service/internal/server"
)

func RegisterRoute(s server.IServer, h IPingHanlder) {
	s.Echo().GET("/ping", h.PingHandler)
}
