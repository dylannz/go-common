package brickftp

import (
	"path/filepath"
	"os"
	"io"
	"github.com/HomesNZ/go-common/util"
	"github.com/Sirupsen/logrus"
	"strings"
	"archive/zip"
	"fmt"
	"path"
)

const (
	// DownloadDir is the path where files will be downloaded to
	DownloadDir = "temp"
)

// Download downloads a file from BrickFTP and saves it in a temp file.
func Download(path string) (string, error) {
	client, err := BrickFTPClient.SFTPConn()
	if err != nil {
		return "", err
	}
	fmt.Println("\t\tCONNECTED")

	src, err := client.Open(path)
	if err != nil {
		logrus.WithField("brickftp", "download").Error(err)
		return "", err
	}
	defer src.Close()

	destPath := filepath.Join(DownloadDir, path)

	// Ensure that the destination directory is created in the temp directory
	util.MustMkdirAll(filepath.Dir(destPath))

	// don't close the destination file because we need to return an open pointer for access to the file. Perhaps this
	// should only return the path to the file that was created and leave it up to the caller to open the file, this way
	// we know that there isn't open files not being GCed.
	dest, err := os.Create(destPath)
	if err != nil {
		logrus.WithField("brickftp", "creating destination file").Error(err)
		return "", err
	}
	defer dest.Close()

	_, err = io.Copy(dest, src)
	if err != nil {
		logrus.WithField("brickftp", "copying to destination").Error(err)
		return "", err
	}

	err = dest.Sync()
	if err != nil {
		logrus.WithField("brickftp", "syncing destination to disk").Error(err)
		return "", err
	}

	return dest.Name(), nil
}

// UnpackZIP unpacks the ZIP file at the given path
func UnpackZIP(dir string) (string, error) {
	wd, err := os.Getwd()
	dir = path.Join(wd, dir)

	reader, err := zip.OpenReader(dir)
	if err != nil {
		logrus.WithField("ZIP", "reading the ZIP archive").Error(err)
		return "", err
	}

	unpackPath := strings.Replace(dir, ".zip", "", 1)

	err = os.MkdirAll(unpackPath, 0755)
	if err != nil {
		logrus.WithField("ZIP", "creating ZIP file directory").Error(err)
		return "", err
	}

	for _, file := range reader.File {
		// ensure that any sub folders are created
		path := filepath.Join(unpackPath, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			logrus.WithField("ZIP", "opening a ZIP file").Error(err)
			return "", err
		}

		targetFile, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, file.Mode())
		if err != nil {
			logrus.WithField("ZIP", "opening ZIP target").Error(err)
			return "", err
		}

		_, err = io.Copy(targetFile, fileReader)
		if err != nil {
			logrus.WithField("ZIP", "copying ZIP file").Error(err)
			return "", err
		}

		fileReader.Close()
		targetFile.Close()
	}

	return unpackPath, nil
}