package go_hydra

import (
	"github.com/FlameInTheDark/go-hydra/connector"
	"github.com/FlameInTheDark/go-hydra/connector/ftp"
	"github.com/FlameInTheDark/go-hydra/connector/http"
	"github.com/FlameInTheDark/go-hydra/connector/imap"
	"github.com/FlameInTheDark/go-hydra/connector/mysql"
	"github.com/FlameInTheDark/go-hydra/connector/mssql"
	"github.com/FlameInTheDark/go-hydra/connector/oracle"
	"github.com/FlameInTheDark/go-hydra/connector/postgres"
	"github.com/FlameInTheDark/go-hydra/connector/redis"
	"github.com/FlameInTheDark/go-hydra/connector/ssh"
	"github.com/FlameInTheDark/go-hydra/connector/telnet"
	"github.com/pkg/errors"
)

type Hydra struct {
	Protocol     connector.IProtocol
	LoginBase    []string
	PasswordBase []string
	Found        map[string]string
	PassOnly     bool
}

const (
	PROTOCOL_FTP    = "ftp"
	PROTOCOL_TELNET = "telnet"
	PROTOCOL_SSH    = "ssh"
	PROTOCOL_HTTP   = "http"
	PROTOCOL_IMAP   = "imap"
	PROTOCOL_REDIS  = "redis"
	PROTOCOL_MYSQL  = "mysql"
	PROTOCOL_MSSQL  = "mssql"
	PROTOCOL_ORACLE  = "oracle"
	PROTOCOL_POSTGRES  = "postgres"
)

func New(login, password []string, protocol, address, port string) (*Hydra, error) {
	var (
		prot     connector.IProtocol
		passOnly bool
	)
	switch protocol {
	case PROTOCOL_FTP:
		prot = ftp.Create(address, port)
	case PROTOCOL_TELNET:
		prot = telnet.Create(address, port)
	case PROTOCOL_SSH:
		prot = ssh.Create(address, port)
	case PROTOCOL_HTTP:
		prot = http.Create(address)
	case PROTOCOL_IMAP:
		prot = imap.Create(address, port)
	case PROTOCOL_REDIS:
		prot = redis.Create(address, port)
		passOnly = true
	case PROTOCOL_MYSQL:
		prot = mysql.Create(address, port)
	case PROTOCOL_MSSQL:
		prot = mssql.Create(address, port)
	case PROTOCOL_ORACLE:
		prot = oracle.Create(address, port)
	case PROTOCOL_POSTGRES:
		prot = postgres.Create(address, port)
	default:
		return nil, errors.New("protocol not found")
	}

	instance := Hydra{prot, login, password, make(map[string]string), passOnly}
	return &instance, nil
}

func (h *Hydra) Check() {
	if h.Protocol.Try() {
		if h.PassOnly {
			for _, pass := range h.PasswordBase {
				if h.Protocol.Check("", pass) {
					h.Found["no_login"] = pass
				}
			}
		} else {
			for _, login := range h.LoginBase {
				for _, pass := range h.PasswordBase {
					if h.Protocol.Check(login, pass) {
						h.Found[login] = pass
					}
				}
			}
		}
	}
}
