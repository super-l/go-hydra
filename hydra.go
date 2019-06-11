package go_hydra

import (
	"github.com/FlameInTheDark/go-hydra/connectors"
	"github.com/FlameInTheDark/go-hydra/connectors/ftp"
	"github.com/FlameInTheDark/go-hydra/connectors/ssh"
	"github.com/FlameInTheDark/go-hydra/connectors/telnet"
	"github.com/pkg/errors"
)

type Hydra struct {
	Protocol     connectors.IProtocol
	LoginBase    []string
	PasswordBase []string
	Founded      map[string]string
}

const (
	PROTOCOL_FTP = "ftp"
	PROTOCOL_TELNET = "telnet"
	PROTOCOL_SSH = "ssh"
)

func New(login, password []string, protocol, address, port string) (*Hydra, error) {
	var prot connectors.IProtocol
	switch protocol {
	case PROTOCOL_FTP:
		prot = ftp.Create(address, port)
	case PROTOCOL_TELNET:
		prot = telnet.Create(address, port)
	case PROTOCOL_SSH:
		prot = ssh.Create(address, port)
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
					h.Founded[login] = pass
				}
			}
		}
	}
}
