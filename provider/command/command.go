package command

import (
	"github.com/hinha/sometor/provider"
	"github.com/spf13/cobra"
)

type Command struct {
	rootCmd *cobra.Command
}

// Fabricate root command
func Fabricate() *Command {
	return &Command{
		rootCmd: newRoot(),
	}
}

// Execute command line interface
func (c *Command) Execute() error {
	return c.rootCmd.Execute()
}

// InjectCommand inject new command into command list
func (c *Command) InjectCommand(scaffolds ...provider.CommandScaffold) {
	for _, scaffold := range scaffolds {
		// Intended assign this variable
		scaffoldRunFunction := scaffold.Run

		cmd := &cobra.Command{
			Use:     scaffold.Use(),
			Short:   scaffold.Short(),
			Example: scaffold.Example(),
			Run: func(cmd *cobra.Command, args []string) {
				scaffoldRunFunction(args)
			},
		}
		c.rootCmd.AddCommand(cmd)
	}
}
