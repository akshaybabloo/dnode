package pkg

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListDirStat(t *testing.T) {
	// Create a temporary directory structure
	tmpDir, err := ioutil.TempDir("", "test-list-dir-stat-*")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	nodeModules1 := filepath.Join(tmpDir, "project1", "node_modules")
	nodeModules2 := filepath.Join(tmpDir, "project2", "node_modules")

	err = os.MkdirAll(nodeModules1, 0755)
	assert.NoError(t, err)
	err = os.MkdirAll(nodeModules2, 0755)
	assert.NoError(t, err)

	// Call ListDirStat
	directories, err := ListDirStat("node_modules", tmpDir)
	assert.NoError(t, err)

	// Check the results
	assert.Len(t, directories, 2)

	for _, dir := range directories {
		switch dir.Path {
		case nodeModules1:
			assert.Equal(t, int64(0), dir.Size)
		case nodeModules2:
			assert.Equal(t, int64(0), dir.Size)
		default:
			t.Fatalf("Unexpected directory path: %s", dir.Path)
		}
	}
}
