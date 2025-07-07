package main

import (
	"github.com/sajad-dev/gingo/cmd/cli"
	"github.com/sajad-dev/gingo/internal/config"
)

func main() {
	config.BootConfig(".env")
	cli.Cli()
}
