package cmd

import (
	"service/internal/app"
	"service/internal/infrastructure"
	"service/internal/migrations"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cli := &cobra.Command{}

	// Add start command
	cli.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Starting server",
		Long:  "Starting server",
		Run: func(cmd *cobra.Command, _ []string) {
			app := app.NewApp()
			e := infrastructure.CreateRoute(app.DB)
			e.Logger.Fatal(e.Start(":8080"))
		},
	})

	// Add migrate up command
	cli.AddCommand(&cobra.Command{
		Use:   "migrate-up",
		Short: "Migrate the schema up",
		Long:  "Migrate the schema up",
		Run: func(cmd *cobra.Command, _ []string) {
			err := migrations.MigrateUp()
			if err != nil {
				panic(err)
			}
		},
	})

	// Add migrate down command
	cli.AddCommand(&cobra.Command{
		Use:   "migrate-down",
		Short: "Migrate the schema down",
		Long:  "Migrate the schema down",
		Run: func(cmd *cobra.Command, _ []string) {
			err := migrations.MigrateDown()
			if err != nil {
				panic(err)
			}
		},
	})

	return cli
}
