package cmd

import "github.com/spf13/cobra"

func NewDeleteCmd() *cobra.Command {
	var deleteCmd = &cobra.Command{
		Use: "delete",
		Short: "Delete all 'node_modules' folders",
	}

	return deleteCmd
}
