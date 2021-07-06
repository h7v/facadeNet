package set_server

import "fmt"

type SetServer interface {
	SetServer()
}

type setServer struct {
	setServer bool
}

//Setting server
func (s *setServer) SetServer() {
	fmt.Println("set server")
}

// Creating NewSetServer
func NewSetServer() *setServer {
	return &setServer{}
}
