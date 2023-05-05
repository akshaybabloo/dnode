package pkg

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Directories holds the Path and Size of a given directory
type Directories struct {
	Path            string    // Absolute Path of a folder
	Size            int64     // Size in bytes
	CreationTime    time.Time // Time when the directory was created
	LastModified    time.Time // Time when the directory was last modified
	NumberOfFiles   int       // Number of files within the directory
	NumberOfSubdirs int       // Number of subdirectories within the directory
}

// ListDirStat lists all the directories starting with the keyword in a given dirPath
// and shows its absolute Directories.Path and total Directories.Size in bytes
func ListDirStat(keyword string, dirPath string) ([]Directories, error) {
	pathStat, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}

	if !pathStat.IsDir() {
		return nil, errors.New("the path provided is not a directory")
	}

	dirChan := make(chan Directories)
	errChan := make(chan error, 1)
	var directories []Directories
	var mu sync.Mutex

	visit := func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() && d.Name() == keyword {
			wg.Add(1)

			go func(path string) {
				defer wg.Done()
				dirStat, err := getTreeSize(path)
				if err != nil {
					errChan <- err
					return
				}
				dirChan <- dirStat
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
		wg.Wait()
		close(dirChan)
		close(errChan)
	}()

	for {
		select {
		case dirStat, ok := <-dirChan:
			if ok {
				mu.Lock()
				directories = append(directories, dirStat)
				mu.Unlock()
			} else {
				dirChan = nil
			}
		case err, ok := <-errChan:
			if ok {
				// You can log the error or decide what to do with it
				fmt.Printf("Error while processing directory: %v\n", err)
			} else {
				errChan = nil
			}
		}

		if dirChan == nil && errChan == nil {
			break
		}
	}

	return directories, nil
}

func getTreeSize(path string) (Directories, error) {
	var totalSize int64
	var numberOfFiles int
	var numberOfSubdirs int
	var creationTime time.Time
	var lastModified time.Time

	err := filepath.WalkDir(path, func(_ string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		if d.IsDir() {
			numberOfSubdirs++
		} else {
			totalSize += info.Size()
			numberOfFiles++
		}

		if creationTime.IsZero() || info.ModTime().Before(creationTime) {
			creationTime = info.ModTime()
		}

		if lastModified.IsZero() || info.ModTime().After(lastModified) {
			lastModified = info.ModTime()
		}

		return nil
	})

	if err != nil {
		return Directories{}, err
	}

	return Directories{
		Path:            path,
		Size:            totalSize,
		CreationTime:    creationTime,
		LastModified:    lastModified,
		NumberOfFiles:   numberOfFiles,
		NumberOfSubdirs: numberOfSubdirs,
	}, nil
}
