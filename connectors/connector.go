package connectors

type IProtocol interface {
	Check(login, password string) bool
	Connect(address, port string) bool
}