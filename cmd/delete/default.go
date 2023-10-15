package delete

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cobra"

	"github.com/akshaybabloo/dnode/pkg"
	walk "github.com/akshaybabloo/go-walk"
)

var wd string
var yes bool

func askForConfirmation() bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(color.GreenString("Are you sure? (y/n): "))
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.ToLower(strings.TrimSpace(input))
		switch input {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("Invalid input. Please try again.")
		}
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
			dirStats, err := walk.ListDirStat(wd, "node_modules")
			if err != nil {
				return err
			}

			if len(dirStats) == 0 {
				s.FinalMSG = color.RedString("No 'node_modules' found")
				s.Stop()
				return nil
			}
			s.Stop()

			totalSize := pkg.PrintDirStats(dirStats, wd)
			fmt.Println()

			if !yes {
				answer := askForConfirmation()
				if !answer {
					return nil
				}
			}
			fmt.Println()

			s.Start()
			var multiErr *multierror.Error
			for _, stat := range dirStats {
				s.Suffix = color.GreenString(" Deleting %s", strings.ReplaceAll(stat.Path, wd, "."))
				if err := os.RemoveAll(stat.Path); err != nil {
					multiErr = multierror.Append(multiErr, err)
				}
			}
			s.FinalMSG = color.GreenString("%s freed", humanize.Bytes(uint64(totalSize)))
			s.Stop()

			return multiErr.ErrorOrNil()
		},
	}

	deleteCmd.Flags().StringVar(&wd, "path", "", "Search path")
	deleteCmd.Flags().BoolVar(&yes, "yes", false, "Delete without confirmation")

	return deleteCmd
}
