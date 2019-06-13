package ssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"strings"
	"time"
)

type SshProtocol struct {
	dst string
}

func (p *SshProtocol) Connect() bool {
	return false
}

func (p *SshProtocol) Try() bool {
	addr := strings.Split(p.dst, ":")
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(addr[0], addr[1]), 5*time.Second)
	if err != nil {
		return false
	}
	tmp := make([]byte, 256)
	if conn != nil {
		_, rErr := conn.Read(tmp)
		if rErr != nil {
			return false
		}
		_ = conn.Close()
	}

	if strings.Contains(strings.ToLower(string(tmp)), "ssh") {
		return true
	} else {
		return false
	}
}

func (p SshProtocol) Check(login, password string) bool {

	conf := &ssh.ClientConfig{
		User: login,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	c, err := ssh.Dial("tcp", p.dst, conf)
	if err != nil {
		return false
	}
	defer c.Close()
	fmt.Printf("[Hydra][ssh] host: %s login: %s password: %s\n", p.dst, login, password)
	return true
}

func Create(address, port string) *SshProtocol {
	var dst string
	if port == "0" {
		dst = address + ":22"
	} else {
		dst = address + ":" + port
	}
	return &SshProtocol{dst: dst}
}
