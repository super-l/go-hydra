package ftp

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

type FtpProtocol struct {
	client *ftp.ServerConn
	dst string
}

func (p *FtpProtocol) Connect() bool {
	c, err := ftp.Dial(p.dst, ftp.DialWithTimeout(1*time.Second))
	if err != nil {
		return false
	}
	p.client = c
	return true
}

func (p *FtpProtocol) Try() bool {
	if p.Connect() {
		_ = p.client.Quit()
		p.client = nil
		return true
	}
	return false
}

func (p FtpProtocol) Check(login, password string) bool {
	p.Connect()
	if p.client != nil {
		err := p.client.Login(login, password)
		defer p.client.Quit()
		if err != nil {
			return false
		}
		fmt.Printf("[Hydra][ftp] host: %s login: %s password: %s\n", p.dst, login, password)
		return true
	}
	return false
}

func Create(address, port string) *FtpProtocol {
	var dst string
	if port == "0" {
		dst = address + ":21"
	} else {
		dst = address + ":" + port
	}
	return &FtpProtocol{dst:dst}
}