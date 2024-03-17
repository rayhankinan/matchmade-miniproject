package cmd

import (
	"service/internal/app"
	"service/internal/migrator"

	"github.com/spf13/cobra"
)

func NewCmd(dep *app.Dep) (cli *cobra.Command) {
	// Instantiate a new cobra command
	cli = &cobra.Command{}

	// Add start command
	cli.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Starting server",
		Long:  "Starting server",
		Run: func(cmd *cobra.Command, _ []string) {
			startCmd := app.NewStartCmd(dep)
			startCmd.StartServer()
		},
	})

	// Add migrate up command
	cli.AddCommand(&cobra.Command{
		Use:   "migrate-up",
		Short: "Migrate the schema up",
		Long:  "Migrate the schema up",
		Run: func(cmd *cobra.Command, _ []string) {
			migrateCmd := migrator.NewMigrateCmd(dep)
			migrateCmd.MigrateUp()
		},
	})

	// Add migrate down command
	cli.AddCommand(&cobra.Command{
		Use:   "migrate-down",
		Short: "Migrate the schema down",
		Long:  "Migrate the schema down",
		Run: func(cmd *cobra.Command, _ []string) {
			migrateCmd := migrator.NewMigrateCmd(dep)
			migrateCmd.MigrateDown()
		},
	})

	return
}
