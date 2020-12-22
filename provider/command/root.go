package command

import "github.com/spf13/cobra"

type root struct{}

func newRoot() *cobra.Command {
	return &cobra.Command{
		Use:     "Api Sometor",
		Short:   "Core business logic of api Sometor",
		Example: "sometor",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
}
