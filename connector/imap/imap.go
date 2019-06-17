package imap

import (
	"fmt"
	"github.com/emersion/go-imap/client"
)

type ImapProtocol struct {
	client *client.Client
	dst    string
}

func (p *ImapProtocol) Connect() bool {
	c, err := client.DialTLS(p.dst, nil)
	if err != nil {
		fmt.Println(err)
		return false
	}
	p.client = c
	return true
}

func (p *ImapProtocol) close() {
	_ = p.client.Close()
	p.client = nil
}

func (p *ImapProtocol) Try() bool {
	if p.Connect() {
		p.close()
		return true
	}
	return false
}

func (p ImapProtocol) Check(login, password string) bool {
	//fmt.Println("Checking for ", login, " : ", password)
	if p.Connect() {
		if err := p.client.Login(login, password); err == nil {
			_ = p.client.Logout()
			p.close()
			fmt.Printf("[Hydra][imap] host: %s login: %s password: %s\n", p.dst, login, password)
			return false
		} else {
			fmt.Println(err)
		}
		p.close()
	}
	return false
}

func Create(address, port string) *ImapProtocol {
	var dst string
	if port == "0" {
		dst = address + ":" + "993"
	} else {
		dst = address + ":" + port
	}
	return &ImapProtocol{dst: dst}
}
