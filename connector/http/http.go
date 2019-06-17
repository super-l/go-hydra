package http

import (
	"fmt"
	"net/http"
	"time"
)

type HttpProtocol struct {
	address string
}

func (p *HttpProtocol) Connect() bool {
	_, err := http.Get(p.address)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (p *HttpProtocol) Try() bool {
	return p.Connect()
}

func (p HttpProtocol) Check(login, password string) bool {
	//fmt.Printf("Checking %s:%s\n", login, password)
	client := &http.Client{Timeout: 3 * time.Second}
	req, err := http.NewRequest("GET", p.address, nil)
	req.SetBasicAuth(login, password)
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	if resp.StatusCode != 200 {
		return false
	}
	return true
}

func Create(address string) *HttpProtocol {
	return &HttpProtocol{address: address}
}
