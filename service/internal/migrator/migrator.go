package migrator

import (
	"service/internal/app"
	"service/internal/config"
	"service/internal/movie"
	"service/internal/user"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IMmigrateCmd interface {
	MigrateUp()
	MigrateDown()
}

type MigrateCmd struct {
	Logger *logrus.Logger
	DB     *gorm.DB
}

func (c *MigrateCmd) MigrateUp() {
	// Migrate the schema
	if err := c.DB.AutoMigrate(&user.User{}, &movie.Movie{}); err != nil {
		// Log the error
		c.Logger.Fatalf("Failed to migrate the schema: %v", err)
		return
	}

	// Log the success
	c.Logger.Info("Schema migration successful")
}

func (c *MigrateCmd) MigrateDown() {
	// Drop the schema
	if err := c.DB.Migrator().DropTable(&movie.Movie{}, &user.User{}); err != nil {
		// Log the error
		c.Logger.Fatalf("Failed to drop the schema: %v", err)
		return
	}

	// Log the success
	c.Logger.Info("Schema drop successful")
}

func NewMigrateCmd(dep *app.Dep) (cmd IMmigrateCmd) {
	// Create a new configuration instance
	cfg, err := config.LoadEnvironment()
	if err != nil {
		dep.Logger.Fatalf("Failed to load environment configuration: %v", err)
	}

	// Create a new database connection
	db, err := app.NewDatabaseConn(cfg)
	if err != nil {
		dep.Logger.Fatalf("Failed to connect to database: %v", err)
	}

	cmd = &MigrateCmd{
		Logger: dep.Logger,
		DB:     db,
	}
	return
}
