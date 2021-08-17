package delete

import (
	"fmt"
	"os"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/akshaybabloo/dnode/pkg"
)

var wd string

// NewDeleteCmd command function to delete node_modules folders
func NewDeleteCmd() *cobra.Command {
	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete all 'node_modules' folders",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			if wd == "" {
				wd, err = os.Getwd()
				if err != nil {
					return err
				}
			}

			fmt.Printf("\r%s", color.GreenString("Searching..."))
			dirStats, err := pkg.ListDirStat("node_modules", wd)
			if err != nil {
				return err
			}

			if len(dirStats) == 0 {
				fmt.Printf("\r")
				color.Red("No 'node_modules' found")
				return nil
			}
			fmt.Printf("\r")

			var totalSize int64 = 0

			for _, stat := range dirStats {
				err := os.RemoveAll(stat.Path)
				if err != nil {
					return err
				}
				fmt.Printf("Deleted %s\n", strings.ReplaceAll(stat.Path, wd, "."))
				totalSize += stat.Size
			}

			color.Green("%s freed", humanize.Bytes(uint64(totalSize)))

			return nil
		},
	}

	deleteCmd.Flags().StringVar(&wd, "path", "", "Search path")

	return deleteCmd
}
