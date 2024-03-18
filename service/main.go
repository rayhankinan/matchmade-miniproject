package main

import (
	"service/internal/cmd"
)

func main() {
	cli := cmd.Command()
	_ = cli.Execute()
}
