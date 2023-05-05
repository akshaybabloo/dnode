package list

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewListCmd(t *testing.T) {
	// Create a temporary directory structure
	tmpDir, err := os.MkdirTemp("", "test-list-cmd-*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	nodeModules1 := filepath.Join(tmpDir, "project1", "node_modules")
	nodeModules2 := filepath.Join(tmpDir, "project2", "node_modules")

	err = os.MkdirAll(nodeModules1, 0755)
	require.NoError(t, err)
	err = os.MkdirAll(nodeModules2, 0755)
	require.NoError(t, err)

	// Set up the Cobra command and flags
	listCmd := NewListCmd()
	wd = tmpDir

	// Redirect os.Stdout to a pipe
	oldStdout := os.Stdout
	readPipe, writePipe, _ := os.Pipe()
	os.Stdout = writePipe
	require.NoError(t, err)

	// Call Execute
	err = listCmd.Execute()
	assert.NoError(t, err)

	// Close the writer and read the output from the reader
	var buf bytes.Buffer
	writePipe.Close()
	io.Copy(&buf, readPipe)
	output := buf.String()

	// Restore os.Stdout
	os.Stdout = oldStdout

	// Check if the output contains the specified substrings
	assert.Contains(t, output, "PATH")
	assert.Contains(t, output, "DIRECTORY SIZE")
	assert.Contains(t, output, "0 B")
	assert.Contains(t, output, "TOTAL")
}
