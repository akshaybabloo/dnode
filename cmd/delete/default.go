package deletecmd

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

const (
	spinnerCharSet = 11
	spinnerDelay   = 100 * time.Millisecond
)

var workingDir string
var confirmDeletion bool

func askForConfirmation() bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(color.GreenString("Are you sure? (y/n): "))
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("error reading input:", err)
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

// NewDeleteCmd creates a new command to delete node_modules folders
func NewDeleteCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete all 'node_modules' folders",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			if workingDir == "" {
				workingDir, err = os.Getwd()
				if err != nil {
					return err
				}
			}

			s := spinner.New(spinner.CharSets[spinnerCharSet], spinnerDelay)
			s.Suffix = color.GreenString(" Searching...")
			s.Start()
			dirStats, err := walk.ListDirStat(workingDir, "node_modules")
			if err != nil {
				return err
			}

			if len(dirStats) == 0 {
				s.FinalMSG = color.RedString("No 'node_modules' found\n")
				s.Stop()
				return nil
			}
			s.Stop()

			totalSize := pkg.PrintDirStats(dirStats, workingDir)
			fmt.Println()

			if !confirmDeletion {
				answer := askForConfirmation()
				if !answer {
					return nil
				}
			}
			fmt.Println()

			s.Start()
			var multiErr *multierror.Error
			for _, dirStat := range dirStats {
				s.Suffix = color.GreenString(" Deleting %s", strings.ReplaceAll(dirStat.Path, workingDir, "."))
				if err := os.RemoveAll(dirStat.Path); err != nil {
					multiErr = multierror.Append(multiErr, err)
				}
			}
			s.FinalMSG = color.GreenString("%s freed", humanize.Bytes(uint64(totalSize)))
			s.Stop()

			if multiErr != nil {
				return multiErr.ErrorOrNil()
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&workingDir, "path", "", "Path to search for 'node_modules' folders")
	cmd.Flags().BoolVar(&confirmDeletion, "yes", false, "Automatically confirm deletion without prompting")

	return cmd
}
