package delete

import (
	"fmt"
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
var yes bool

func askForConfirmation() bool {
	var response string
	fmt.Print(color.GreenString("Are you sure? (y/n): "))
	// TODO: user readstring("\n") instead of scanln
	_, err := fmt.Scanln(&response)
	if err != nil {
		// empty new line should be ignored
		return askForConfirmation()
	}

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		fmt.Println("I'm sorry but I didn't get what you meant, please type (y)es or (n)o and then press enter:")
		return askForConfirmation()
	}
}

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
			fmt.Println()

			if !yes {
				answer := askForConfirmation()
				if !answer {
					return nil
				}
			}
			fmt.Println()

			s.Start()
			for _, stat := range dirStats {
				err := os.RemoveAll(stat.Path)
				if err != nil {
					return err
				}
				s.Suffix = color.GreenString("Deleted %s", strings.ReplaceAll(stat.Path, wd, "."))
			}
			s.Stop()

			color.Green("%s freed", humanize.Bytes(uint64(totalSize)))

			return nil
		},
	}

	deleteCmd.Flags().StringVar(&wd, "path", "", "Search path")
	deleteCmd.Flags().BoolVar(&yes, "yes", false, "Delete without confirmation")

	return deleteCmd
}
