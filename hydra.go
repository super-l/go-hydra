package go_hydra

import (
	"github.com/FlameInTheDark/go-hydra/connector"
	"github.com/FlameInTheDark/go-hydra/connector/ftp"
	"github.com/FlameInTheDark/go-hydra/connector/http"
	"github.com/FlameInTheDark/go-hydra/connector/imap"
	"github.com/FlameInTheDark/go-hydra/connector/ssh"
	"github.com/FlameInTheDark/go-hydra/connector/telnet"
	"github.com/pkg/errors"
)

type Hydra struct {
	Protocol     connector.IProtocol
	LoginBase    []string
	PasswordBase []string
	Found        map[string]string
}

const (
	PROTOCOL_FTP    = "ftp"
	PROTOCOL_TELNET = "telnet"
	PROTOCOL_SSH    = "ssh"
	PROTOCOL_HTTP   = "http"
	PROTOCOL_IMAP   = "imap"
)

func New(login, password []string, protocol, address, port string) (*Hydra, error) {
	var prot connector.IProtocol
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
	default:
		return nil, errors.New("protocol not found")
	}

	instance := Hydra{prot, login, password, make(map[string]string)}
	return &instance, nil
}

func (h *Hydra) Check() {
	if h.Protocol.Try() {
		for _, login := range h.LoginBase {
			for _, pass := range h.PasswordBase {
				if h.Protocol.Check(login, pass) {
					h.Found[login] = pass
				}
			}
		}
	}
}
