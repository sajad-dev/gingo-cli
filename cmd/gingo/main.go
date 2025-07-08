package main

import (
	"github.com/sajad-dev/gingo-cli/internal/cli"
	"github.com/sajad-dev/gingo-cli/internal/config"
)

func main() {
	config.BootConfig(".env")
	cli.Cli()
}
