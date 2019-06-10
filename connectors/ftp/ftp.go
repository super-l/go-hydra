package ftp

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

type FtpProtocol struct {
	client *ftp.ServerConn
	address string
	port string
}

func (p *FtpProtocol) Connect() bool {
	c, err := ftp.Dial(p.address+":"+p.port, ftp.DialWithTimeout(1*time.Second))
	if err != nil {
		fmt.Println(err)
		return false
	}
	p.client = c
	return true
}

func (p *FtpProtocol) Check(login, password string) bool {
	p.Connect()
	if p.client != nil {
		err := p.client.Login(login, password)
		defer p.client.Quit()
		if err != nil {
			return false
		}
		fmt.Printf("[ftp] host: %s:%s login: %s password: %s\n", p.address, p.port, login, password)
		return true
	}
	return false
}

func Create(address, port string) *FtpProtocol {
	return &FtpProtocol{address:address, port:port}
}