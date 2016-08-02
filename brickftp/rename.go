package brickftp

import "github.com/Sirupsen/logrus"

// Rename renames a file on the BrickFTP server.
func Rename(oldname string, newname string) error {
	client, err := BrickFTPClient.SFTPConn()
	if err != nil {
		return err
	}

	err = client.Rename(oldname, newname)
	if err != nil {
		logrus.WithField("brickftp", "rename").Error(err)
		return err
	}
	return nil
}
