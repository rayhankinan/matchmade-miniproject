package healthcheck

import (
	"service/internal/server"
)

func RegisterRoute(s server.IServer, h IHealthcheckHandler) {
	s.Echo().GET("/ping", h.PingHandler)
}
