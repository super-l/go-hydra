package oracle

import (
	"database/sql"
	"fmt"
	_ "gopkg.in/goracle.v2"
)

// needs: import _ "gopkg.in/goracle.v2"
// or: import _ "gopkg.in/rana/ora.v4"

type OracleProtocol struct {
	dst string
}

func (p *OracleProtocol) Connect() bool {
	return true
}

func (p *OracleProtocol) Try() bool {
	return true
}

func (p OracleProtocol) Check(login, password string) bool {
	db, err := sql.Open("goracle",fmt.Sprintf("%s/%s@%s/", login, password, p.dst))
	if err != nil {
		return false
	}
	dbErr := db.Ping()
	if dbErr != nil {
		return false
	}
	fmt.Printf("[Hydra][oracle] host: %s login: %s password: %s\n", p.dst, login, password)
	return true
}

func Create(address, port string) *OracleProtocol {
	var dst string
	if port == "0" {
		dst = address + ":1521"
	} else {
		dst = address + ":" + port
	}
	return &OracleProtocol{dst}
}
