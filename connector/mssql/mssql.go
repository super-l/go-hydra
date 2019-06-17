package mssql

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"net/url"
)

// needs: import _ "github.com/denisenkom/go-mssqldb"

type MssqlProtocol struct {
	dst string
}

func (p *MssqlProtocol) Connect() bool {
	return true
}

func (p *MssqlProtocol) Try() bool {
	return true
}

func (p MssqlProtocol) Check(login, password string) bool {
	query := url.Values{}
	query.Add("app name", "hydra")
	u := url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(login, password),
		Host:     p.dst,
		RawQuery: query.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		return false
	}
	dbErr := db.Ping()
	if dbErr != nil {
		return false
	}
	fmt.Printf("[Hydra][mssql] host: %s login: %s password: %s\n", p.dst, login, password)
	return true
}

func Create(address, port string) *MssqlProtocol {
	var dst string
	if port == "0" {
		dst = address + ":1433"
	} else {
		dst = address + ":" + port
	}
	return &MssqlProtocol{dst}
}
