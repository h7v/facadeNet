package set_server

import "fmt"

// Set server interface
type SetServer interface {
	SetServer()
}

type setServer struct {
	setServer bool
}

// Check how server was set
func (s *setServer) CheckSetServer() string {
	return "set server"
}

// Setting server
func (s *setServer) SetServer() {
	fmt.Println("set server")
}

// Creating NewSetServer
func NewSetServer() *setServer {
	return &setServer{}
}
