package list

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/akshaybabloo/dnode/pkg"
)

var wd string

// NewListCmd command function to list all node_modules folders
func NewListCmd() *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists out all 'node_modules' folders",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			if wd == "" {
				wd, err = os.Getwd()
				if err != nil {
					return err
				}
			}

			s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
			s.Suffix = color.GreenString(" Searching...")
			s.Start()
			dirStats, err := pkg.ListDirStat("node_modules", wd)
			if err != nil {
				return err
			}

			if len(dirStats) == 0 {
				s.FinalMSG = color.RedString("No 'node_modules' found")
				s.Stop()
				return nil
			}
			s.Stop()

			_ = pkg.PrintDirStats(dirStats, wd)

			return nil
		},
	}

	listCmd.Flags().StringVar(&wd, "path", "", "Search path")

	return listCmd
}
