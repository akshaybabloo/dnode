package delete

import "github.com/spf13/cobra"

func NewDeleteCmd() *cobra.Command {
	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete all 'node_modules' folders",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	deleteCmd.Flags().StringVar(&wd, "path", "", "Search path")

	return deleteCmd
}
