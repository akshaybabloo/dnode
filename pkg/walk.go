package pkg

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

type Directories struct {
	Path string
	Size int
}

// ListDirStat lists all the directories starting with the keyword in a given dirPath
func ListDirStat(keyword string, dirPath string) ([]Directories, error) {
	pathStat, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}

	if !pathStat.IsDir() {
		return nil, errors.New("the path provided is not a directory")
	}

	err = filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {

		return err
	})
	if err != nil {
		return nil, err
	}

	return []Directories{}, nil
}
