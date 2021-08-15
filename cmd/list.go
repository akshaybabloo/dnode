package cmd

import "github.com/spf13/cobra"

func NewListCmd() *cobra.Command {
	var listCmd = &cobra.Command{
		Use: "list",
		Short: "Lists out all 'node_modules' folders",
	}

	return listCmd
}
