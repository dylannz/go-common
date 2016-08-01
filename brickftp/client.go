package brickftp

import (
	"errors"
	"sync"

	"github.com/HomesNZ/go-common/env"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var (
	// BrickFTPClient is a brickftp.Client singleton.
	BrickFTPClient *Client
)

// InitClient initializes BrickFTPClient with the default brickftp.Client
func InitClient() {
	BrickFTPClient = &Client{}
}

// Client is a BrickFTP client consisting of a SSH client and SFTP client, which utilizes the SSH client connection.
type Client struct {
	sync.Mutex
	sshClient  *ssh.Client
	sftpClient *sftp.Client
}

// sshConn guarantees returning a connected SHH client or an error if a connection cannot be established.
func (c *Client) sshConn() (*ssh.Client, error) {
	// Return early if the ssh client is already connected.
	if c.sshClientConnected() {
		return c.sshClient, nil
	}

	client, err := ssh.Dial("tcp", env.MustGetString("BRICKFTP_ADDRESS"), connectionConfig())
	if err != nil {
		return c.sshClient, err
	}

	c.sshClient = client
	if c.sshClientConnected() {
		return c.sshClient, nil
	}

	return nil, errors.New("Cannot create SSH connection.")
}

func (c *Client) sshClientConnected() bool {
	// Return early if the sshClient has never been created.
	if c.sshClient == nil {
		return false
	}

	// Attempt to create a new session to verify the connection
	s, err := c.sshClient.NewSession()
	if err != nil {
		return false
	}
	s.Close()

	return true
}

// SFTPConn guarantees returning a connected SFTP client or an error if a connection cannot be established.
func (c *Client) SFTPConn() (*sftp.Client, error) {
	// Return early if the ssh client is already connected.
	if c.sftpClientConnected() {
		return c.sftpClient, nil
	}

	sshClient, err := c.sshConn()
	if err != nil {
		return c.sftpClient, err
	}

	client, err := sftp.NewClient(sshClient)
	if err != nil {
		return c.sftpClient, err
	}

	c.sftpClient = client
	if c.sftpClientConnected() {
		return c.sftpClient, nil
	}

	return nil, errors.New("Cannot create SSH connection.")
}

func (c *Client) sftpClientConnected() bool {
	// Return early if the sftpClient has never been created.
	if c.sftpClient == nil {
		return false
	}

	_, err := c.sftpClient.Getwd()
	if err != nil {
		return false
	}
	return true
}

func connectionConfig() *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: env.MustGetString("BRICKFTP_USER"),
		Auth: []ssh.AuthMethod{
			ssh.Password(env.MustGetString("BRICKFTP_PASSWORD")),
		},
	}
}
