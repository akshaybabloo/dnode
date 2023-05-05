package delete

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeleteCmd(t *testing.T) {
	// Create a temporary directory structure
	tmpDir, err := os.MkdirTemp("", "test-delete-cmd-*")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	nodeModules1 := filepath.Join(tmpDir, "project1", "node_modules")
	nodeModules2 := filepath.Join(tmpDir, "project2", "node_modules")

	err = os.MkdirAll(nodeModules1, 0755)
	assert.NoError(t, err)
	err = os.MkdirAll(nodeModules2, 0755)
	assert.NoError(t, err)

	// Set up the Cobra command and flags
	deleteCmd := NewDeleteCmd()
	wd = tmpDir
	yes = true

	// Call Execute
	err = deleteCmd.Execute()
	assert.NoError(t, err)

	// Check if node_modules directories have been deleted
	_, err = os.Stat(nodeModules1)
	assert.True(t, os.IsNotExist(err))

	_, err = os.Stat(nodeModules2)
	assert.True(t, os.IsNotExist(err))
}
