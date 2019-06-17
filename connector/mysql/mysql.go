package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// needs: import _ "github.com/go-sql-driver/mysql"

type MysqlProtocol struct {
	dst string
}

func (p *MysqlProtocol) Connect() bool {
	return true
}

func (p *MysqlProtocol) Try() bool {
	return true
}

func (p MysqlProtocol) Check(login, password string) bool {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/", login, password, p.dst))
	if err != nil {
		return false
	}
	dbErr := db.Ping()
	if dbErr != nil {
		return false
	}
	fmt.Printf("[Hydra][mysql] host: %s login: %s password: %s\n", p.dst, login, password)
	return true
}

func Create(address, port string) *MysqlProtocol {
	var dst string
	if port == "0" {
		dst = address + ":3306"
	} else {
		dst = address + ":" + port
	}
	return &MysqlProtocol{dst}
}
