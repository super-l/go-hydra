package redis

import (
	"fmt"
	"github.com/mediocregopher/radix"
)

type RedisProtocol struct {
	dst string
}

func (p *RedisProtocol) Connect() bool {
	return true
}

func (p *RedisProtocol) Try() bool {
	return true
}

func (p RedisProtocol) Check(login, password string) bool {
	c, err := radix.Dial("tcp", p.dst, radix.DialAuthPass(password))

	if err != nil {
		//fmt.Println(err)
		return false
	}
	var resp string
	dErr := c.Do(radix.Cmd(&resp, "PING"))
	if dErr != nil {
		//fmt.Println(dErr)
		return false
	}
	if resp == "PONG" {
		//fmt.Println(resp)
		fmt.Printf("[Hydra][redis] host: %s password: %s\n", p.dst, password)
		return true
	}
	return false
}

func Create(address, port string) *RedisProtocol {
	var dst string
	if port == "0" {
		dst = address + ":6379"
	} else {
		dst = address + ":" + port
	}
	return &RedisProtocol{dst: dst}
}
