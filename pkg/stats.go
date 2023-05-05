package pkg

import (
	"os"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func PrintDirStats(dirStats []Directories, wd string) int64 {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Path", "Directory Size"})
	var totalSize int64
	for _, stat := range dirStats {
		t.AppendRow([]interface{}{
			strings.ReplaceAll(stat.Path, wd, "."),
			humanize.Bytes(uint64(stat.Size)),
		})
		totalSize += stat.Size
	}
	t.SetColumnConfigs([]table.ColumnConfig{
		{Name: "Path", Align: text.AlignLeft, AlignHeader: text.AlignCenter},
		{Name: "Directory Size", Align: text.AlignRight, AlignHeader: text.AlignCenter},
	})
	t.AppendFooter(table.Row{"Total", humanize.Bytes(uint64(totalSize))})
	t.Render()

	return totalSize
}
