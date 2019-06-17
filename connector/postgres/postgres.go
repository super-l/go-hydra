package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// needs: import _ "github.com/lib/pq"

type PostgresProtocol struct {
	dst string
}

func (p *PostgresProtocol) Connect() bool {
	return true
}

func (p *PostgresProtocol) Try() bool {
	return true
}

func (p PostgresProtocol) Check(login, password string) bool {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/?sslmode=verify-full", login, password, p.dst))
	if err != nil {
		return false
	}
	dbErr := db.Ping()
	if dbErr != nil {
		return false
	}
	fmt.Printf("[Hydra][postgres] host: %s login: %s password: %s\n", p.dst, login, password)
	return true
}

func Create(address, port string) *PostgresProtocol {
	var dst string
	if port == "0" {
		dst = address + ":5432"
	} else {
		dst = address + ":" + port
	}
	return &PostgresProtocol{dst}
}
