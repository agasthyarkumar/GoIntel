package compression

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Compress(source string) error {

	// Clean path
	source = filepath.Clean(source)

	// Parent directory
	parentDir := filepath.Dir(source)

	// Base filename
	baseName := filepath.Base(source)

	// Final zip path
	zipPath := filepath.Join(parentDir, baseName+".zip")

	// Create zip file
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	info, err := os.Stat(source)
	if err != nil {
		return err
	}

	// Single file
	if !info.IsDir() {
		return addFile(zipWriter, source)
	}

	// Directory
	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		return addFile(zipWriter, path)
	})
}

func addFile(zipWriter *zip.Writer, path string) error {

	fileToZip, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = path

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)

	return err
}