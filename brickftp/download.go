package brickftp

import (
	"path/filepath"
	"os"
	"io"
	"github.com/HomesNZ/go-common/util"
	"github.com/Sirupsen/logrus"
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
