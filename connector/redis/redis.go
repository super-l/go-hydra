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
	c, err := radix.Dial("tcp", "localhost:6379", radix.DialAuthPass(password))
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
	//
	//client := redis.NewClient(&redis.Options{
	//	Addr:     p.dst,
	//	Password: password,
	//	DB:       0, // use default DB
	//})
	//defer client.Close()
	//
	//pong, err := client.Ping().Result()
	//if err != nil {
	//	fmt.Println(err)
	//	return false
	//}
	//fmt.Println(pong)
	//fmt.Printf("[Hydra][redis] host: %s password: %s\n", p.dst, password)
	//return true
}

func Create(address, port string) *RedisProtocol {
	var dst string
	if port == "0" {
		dst = address + ":3679"
	} else {
		dst = address + ":" + port
	}
	return &RedisProtocol{dst: dst}
}
