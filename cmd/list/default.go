package list

import (
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
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

			var totalSize int64 = 0

			table := tablewriter.NewWriter(os.Stdout)
			for _, stat := range dirStats {
				table.Append([]string{strings.ReplaceAll(stat.Path, wd, "."), humanize.Bytes(uint64(stat.Size))})
				totalSize += stat.Size
			}
			table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_CENTER})
			table.SetHeader([]string{"Path", "Directory Size"})
			table.SetFooter([]string{"Total", humanize.Bytes(uint64(totalSize))})
			table.SetAutoMergeCellsByColumnIndex([]int{2, 3})
			table.SetBorder(false)
			table.Render()

			return nil
		},
	}

	listCmd.Flags().StringVar(&wd, "path", "", "Search path")

	return listCmd
}
