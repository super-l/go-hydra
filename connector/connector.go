package connector

type IProtocol interface {
	// Check login password
	Check(login, password string) bool
	// Try checks connection to host
	Try() bool
	// Connect to host
	Connect() bool
}