package server

import (
	"context"
	"fmt"

	"service/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type IServer interface {
	Start(config.FlagConfig) error
	Stop(ctx context.Context) error
	Echo() *echo.Echo
}

type Server struct {
	echo   *echo.Echo
	logger *logrus.Logger
}

func (s *Server) Start(cfg config.FlagConfig) (err error) {
	err = s.echo.Start(fmt.Sprintf(":%d", cfg.Port))
	return
}

func (s *Server) Stop(ctx context.Context) (err error) {
	err = s.echo.Shutdown(ctx)
	return
}

func (s *Server) Echo() (e *echo.Echo) {
	e = s.echo
	return
}

func NewServer(l *logrus.Logger) (server IServer) {
	// Create a new echo instance
	e := echo.New()
	e.Logger.SetOutput(l.Writer())

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Create a new server instance
	server = &Server{
		echo:   e,
		logger: l,
	}

	return
}
