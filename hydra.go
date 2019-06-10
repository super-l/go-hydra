package go_hydra

import (
	"github.com/FlameInTheDark/go-hydra/connectors"
	"github.com/FlameInTheDark/go-hydra/connectors/ftp"
)

type Hydra struct {
	Protocol     connectors.IProtocol
	LoginBase    []string
	PasswordBase []string
	Founded      map[string]string
}

const (
	PROTOCOL_FTP = "ftp"
)

func New(login, password []string, protocol, address, port string) (*Hydra, error) {
	var prot connectors.IProtocol
	switch protocol {
	case PROTOCOL_FTP:
		prot = ftp.Create(address, port)
	}
	instance := Hydra{prot, login, password, make(map[string]string)}
	return &instance, nil
}

func (h *Hydra) Check() {
	for _, login := range h.LoginBase {
		for _, pass := range h.PasswordBase {
			if h.Protocol.Check(login, pass) {
				h.Founded[login] = pass
			}
		}
	}
}
