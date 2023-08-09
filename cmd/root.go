package cmd

import "github.com/spf13/cobra"

func Root() *cobra.Command {
	root := cobra.Command{}

	root.AddCommand(server())

	return &root
}
