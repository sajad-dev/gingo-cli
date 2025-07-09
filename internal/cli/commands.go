package cli

import (
	"os"

	"github.com/sajad-dev/gingo-cli/internal/command/newpr"
	"github.com/sajad-dev/gingo-cli/internal/command/versions"
	"github.com/spf13/cobra"
)

var CommandsList []*Commands = []*Commands{
	{
		Command: &cobra.Command{
			Use:   "new",
			Short: "initialize new module in current directory",
			Run: func(cmd *cobra.Command, args []string) {
				if err := newpr.Init(cmd, args); err != nil {
					os.Exit(0)
				}
			},
		},
		Flags: []any{
			FlagString{Flag: "repo", ShortFlag: "r", Defualt: "gingo-project", Discription: "The Git repository URL to use as a template. If omitted, the default template will be used."},
			FlagString{Flag: "name", ShortFlag: "n", Defualt: "gingo-project", Discription: "The name of the new project directory."},
			FlagString{Flag: "version", ShortFlag: "v", Defualt: "latest", Discription: "The version of the template or framework to install (e.g., v1.2.3 or latest)."},
		},
	},
	{
		Command: &cobra.Command{
			Use:   "version",
			Short: "Displays the current version of the framework CLI.",
			Run: func(cmd *cobra.Command, args []string) {
				if err := versions.Versions(); err != nil {
					os.Exit(1)
				}
			},
		},
		Flags: []any{},
	},

}

type FlagString struct {
	Value       string
	Flag        string
	ShortFlag   string
	Defualt     string
	Discription string
}
type FlagInt struct {
	Value       int
	Flag        string
	ShortFlag   string
	Defualt     int
	Discription string
}
type Commands struct {
	Command *cobra.Command
	Flags   []any
}
