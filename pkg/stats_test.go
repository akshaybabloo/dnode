package pkg

import (
	"bytes"
	"io"
	"os"
	"testing"

	walk "github.com/akshaybabloo/go-walk"
	"github.com/stretchr/testify/assert"
)

func TestPrintDirStats(t *testing.T) {
	// Create a sample Directories slice
	dirStats := []walk.DirectoryInfo{
		{Path: "/tmp/test1/node_modules", Size: 1024},
		{Path: "/tmp/test2/node_modules", Size: 2048},
	}

	// Redirect os.Stdout to a pipe
	oldStdout := os.Stdout
	readPipe, writePipe, _ := os.Pipe()
	os.Stdout = writePipe

	// Call PrintDirStats
	totalSize := PrintDirStats(dirStats, "/tmp")

	// Check the total size
	assert.Equal(t, int64(3072), totalSize)

	// Capture the output
	var buf bytes.Buffer
	writePipe.Close()
	io.Copy(&buf, readPipe)
	output := buf.String()

	// Restore os.Stdout
	os.Stdout = oldStdout

	// Check if the table has the correct data
	assert.Contains(t, output, "test1/node_modules")
	assert.Contains(t, output, "test2/node_modules")
	assert.Contains(t, output, "1.0 kB") // 1024 bytes
	assert.Contains(t, output, "2.0 kB") // 2048 bytes
	assert.Contains(t, output, "3.1 KB") // Total size

	// Check header alignment
	assert.Contains(t, output, "  PATH   ")
	assert.Contains(t, output, " DIRECTORY SIZE ")
}
