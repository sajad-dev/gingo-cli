package cli

import (
	"os"

	"github.com/fatih/color"
	"github.com/sajad-dev/gingo/internal/commad/newpr"
	"github.com/sajad-dev/gingo/internal/commad/versions"
	"github.com/spf13/cobra"
)

var CommandsList []*Commands = []*Commands{
	{
		Command: &cobra.Command{
			Use:   "new",
			Short: "initialize new module in current directory",
			Run: func(cmd *cobra.Command, args []string) {
				if err := newpr.Init(); err != nil {
					os.Exit(1)
				}
				color.Green("Sucssfuly create project :) - ;)")
			},
		},
		Flags: []any{
			FlagString{Flag: "repo", ShortFlag: "r", Defualt: "", Discription: "The Git repository URL to use as a template. If omitted, the default template will be used."},
			FlagString{Flag: "name", ShortFlag: "n", Defualt: "", Discription: "The name of the new project directory."},
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
	{
		Command: &cobra.Command{
			Use:   "install [modules...]",
			Short: "Install one or more modules",
			Run: func(cmd *cobra.Command, args []string) {
				if err := newpr.Init(); err != nil {
					os.Exit(1)
				}
			},
		},
		Flags: []any{
			FlagString{Flag: "version", ShortFlag: "v", Defualt: "latest", Discription: "The version of the template or framework to install (e.g., v1.2.3 or latest)."},
		},
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
