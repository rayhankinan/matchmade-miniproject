package main

import (
	"service/internal/app"
	"service/internal/cmd"
)

func main() {
	dep := app.NewDep()
	cli := cmd.NewCmd(dep)
	_ = cli.Execute()
}
