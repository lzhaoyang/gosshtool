package gosshtool

import (
	"golang.org/x/crypto/ssh"
	"io"
	"log"
)

type PtyInfo struct {
	Term  string
	H     int
	W     int
	Modes ssh.TerminalModes
}

type ReadWriteCloser interface {
	io.Reader
	io.WriteCloser
}

type SSHClientConfig struct {
	Host              string
	User              string
	Password          string
	PrivateKey        string
	DialTimeoutSecond int
	MaxDataThroughput uint64
}

func makeConfig(user string, password string, privateKey string) (config *ssh.ClientConfig) {

	if password == "" && user == "" {
		log.Fatal("No password or private key available")
	}
	config = &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		//fix issue Failed to dial: ssh: must specify HostKeyCallback
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if privateKey != "" {
		signer, err := ssh.ParsePrivateKey([]byte(privateKey))
		if err != nil {
			log.Fatalf("ssh.ParsePrivateKey error:%v", err)
		}
		clientkey := ssh.PublicKeys(signer)
		config = &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				clientkey,
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	}
	return
}
