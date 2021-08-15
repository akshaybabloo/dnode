package pkg

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

var wg sync.WaitGroup

type Directories struct {
	Path string
	Size int64
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
	dirChan := make(chan Directories, 1)
	var directories []Directories

	visit := func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() && d.Name() == keyword {
			fmt.Println(path)
			wg.Add(1)

			go func(path string) {
				size, _ := getTreeSize(path)
				dirChan <- Directories{
					Path: path,
					Size: size,
				}
			}(path)
			return filepath.SkipDir
		}
		return err
	}

	err = filepath.WalkDir(dirPath, visit)
	if err != nil {
		return nil, err
	}

	go func() {
		for i := range dirChan {
			directories = append(directories, i)
			wg.Done()
		}
	}()

	wg.Wait()

	return directories, nil
}

func getTreeSize(path string) (int64, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}
	var total int64
	for _, entry := range entries {
		if entry.IsDir() {
			size, err := getTreeSize(filepath.Join(path, entry.Name()))
			if err != nil {
				return 0, err
			}
			total += size
		} else {
			info, err := entry.Info()
			if err != nil {
				return 0, err
			}
			total += info.Size()
		}
	}
	return total, nil
}
