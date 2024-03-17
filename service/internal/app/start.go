package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"service/internal/config"
	"service/internal/healthcheck"
	"service/internal/server"

	"github.com/sirupsen/logrus"
)

type IStartCmd interface {
	StartServer()
}

type StartCmd struct {
	Server   server.IServer
	Logger   *logrus.Logger
	QuitChan chan os.Signal
}

func (c *StartCmd) StartServer() {
	// Create a new goroutine to start the server
	go func() {
		// Load the configuration
		cfg, err := config.LoadFlags()
		if err != nil {
			c.Logger.Fatalf("Failed to load flags: %v", err)
		}

		// Start the server
		if err := c.Server.Start(cfg); err != nil && errors.Is(err, http.ErrServerClosed) {
			c.Logger.Fatalf("Shutting down the server: %v", err)
		}
	}()

	// Notify the server to shutdown
	signal.Notify(c.QuitChan, os.Interrupt)
	<-c.QuitChan

	c.Logger.Info("Attempting graceful shutdown")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.Server.Stop(ctx); err != nil {
		c.Logger.Fatal(err)
	}

	c.Logger.Info("Graceful shutdown completed")
}

func NewStartCmd(dep *Dep) (cmd IStartCmd) {
	// Create a new server instance
	s := server.NewServer(dep.Logger)

	// Create a new configuration instance
	_, err := config.LoadEnvironment()
	if err != nil {
		dep.Logger.Fatalf("Failed to load environment configuration: %v", err)
	}

	// Create a new database connection
	// db, err = NewDatabaseConn(cfg)
	// if err != nil {
	// 	dep.Logger.Fatalf("Failed to connect to database: %v", err)
	// }

	// Create a new handler and register the route
	healthcheckHandler := healthcheck.NewHandle()
	healthcheck.RegisterRoute(s, healthcheckHandler)

	// Create a new quit channel
	quitChan := make(chan os.Signal, 1)

	// Create a new start command
	cmd = &StartCmd{
		Server:   s,
		Logger:   dep.Logger,
		QuitChan: quitChan,
	}

	return
}
