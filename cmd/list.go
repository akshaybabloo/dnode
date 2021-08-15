package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/akshaybabloo/dnode/pkg"
)

var wd string

func NewListCmd() *cobra.Command {
	var listCmd = &cobra.Command{
		Use: "list",
		Short: "Lists out all 'node_modules' folders",
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

			table := tablewriter.NewWriter(os.Stdout)
			for _, stat := range dirStats {
				table.Append([]string{strings.ReplaceAll(stat.Path, wd, "."), humanize.Bytes(uint64(stat.Size))})
				totalSize += stat.Size
			}
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
