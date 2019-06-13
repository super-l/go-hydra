package telnet

import (
	"fmt"
	"github.com/ziutek/telnet"
	"time"
)

type TelnetProtocol struct {
	client *telnet.Conn
	dst    string
}

func (p *TelnetProtocol) Connect() bool {
	c, err := telnet.Dial("tcp", p.dst)
	if err != nil {
		fmt.Println("Connecting error: ", err)
		return false
	}
	p.client = c
	return true
}

func (p *TelnetProtocol) Try() bool {
	if p.Connect() {
		_ = p.client.Close()
		p.client = nil
		return true
	}
	return false
}

const timeout = 2 * time.Second

func checkErr(err error) {
	if err != nil {
		//fmt.Println("Error:", err)
	}
}

func expect(t *telnet.Conn, d ...string) {
	checkErr(t.SetReadDeadline(time.Now().Add(timeout)))
	checkErr(t.SkipUntil(d...))
}

func sendln(t *telnet.Conn, s string) {
	checkErr(t.SetWriteDeadline(time.Now().Add(timeout)))
	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = '\n'
	_, err := t.Write(buf)
	checkErr(err)
}

func (p TelnetProtocol) Check(login, password string) bool {
	if p.client == nil {
		//fmt.Println("Connecting to ", p.address + ":" + p.port)
		p.Connect()
	}
	if p.client != nil {
		defer p.client.Close()
		defer func() { p.client = nil }()
		//fmt.Printf("Checking: %s %s\n", login, password)
		expect(p.client, "login: ")
		sendln(p.client, login)
		expect(p.client, "ssword: ")
		sendln(p.client, password)
		expect(p.client, "$")
		sendln(p.client, "uname -a")
		data, err := p.client.ReadBytes('$')
		if err != nil {
			//fmt.Println("Reading error: ", err)
			return false
		}
		if len(data) > 0 {
			fmt.Printf("[Hydra][telnet] host: %s login: %s password: %s\n", p.dst, login, password)
			return true
		}
		return false
	}
	return false
}

func Create(address, port string) *TelnetProtocol {
	var dst string
	if port == "0" {
		dst = address + ":23"
	} else {
		dst = address + ":" + port
	}
	return &TelnetProtocol{dst:dst}
}
