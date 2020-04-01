package set_server

import "fmt"

// SetServer interface
type SetServer interface {
	SetServer()
}

type setServer struct {
	setServer bool
}

// CheckSetServer checks how server was set
func (s *setServer) CheckSetServer() (setServerArg string) {
	return "set server"
}

// SetServer is setting server
func (s *setServer) SetServer() {
	fmt.Println("set server")
}

// NewSetServer creating 
func NewSetServer() (setServerArg *setServer) {
	return &setServer{}
}
