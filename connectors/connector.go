package connectors

type IProtocol interface {
	Check(login, password string) bool
	Connect() bool
}