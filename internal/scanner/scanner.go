package scanner

import (
	"os"
	"path/filepath"
)

func ScanDirectory(root string, jobs chan<- string) error {

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		jobs <- path

		return nil
	})

	close(jobs)

	return err
}