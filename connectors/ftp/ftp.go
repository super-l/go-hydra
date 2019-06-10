package ftp

import (
	"github.com/jlaffaye/ftp"
	"time"
)

type FtpProtocol struct {
	client *ftp.ServerConn
}

func (p *FtpProtocol) Connect(address, port string) bool {
	c, err := ftp.Dial(address+":"+port, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return false
	}
	p.client = c
	return true
}

func (p FtpProtocol) Check(login, password string) bool {
	if p.client != nil {
		err := p.client.Login(login, password)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

func Create() *FtpProtocol {
	return &FtpProtocol{}
}