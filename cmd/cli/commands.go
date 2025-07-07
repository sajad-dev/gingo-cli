package cli

import (
	"github.com/spf13/cobra"
)

var CommandsList []*Commands = []*Commands{

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
