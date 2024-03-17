package cmd

import (
	"service/internal/app"

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

	return
}
