package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"

	delete2 "github.com/akshaybabloo/dnode/cmd/delete"
	"github.com/akshaybabloo/dnode/cmd/list"
)

func NewRootCmd(appVersion, buildDate string) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "dnode [OPTIONS] [COMMANDS]",
		Short: "Tool to delete 'node_modules'",
		Long:  `dnode can be used to delete 'node_modules' recursively from sub-folders`,
		Example: heredoc.Doc(`
			$ dnode list
			$ dnode delete <directory path>
			$ dnode delete
		`),
	}

	rootCmd.AddCommand(list.NewListCmd())
	rootCmd.AddCommand(delete2.NewDeleteCmd())
	return rootCmd
}
