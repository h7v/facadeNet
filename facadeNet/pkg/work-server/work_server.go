package work_server

import "fmt"

// WorkServer interface
type WorkServer interface {
	Work()
}

type workServer struct {
	workServer bool
}

// Check how server was work
func (w *workServer) CheckWorkServer() (workServerArg string) {
	return "work server"
}

//Work of the server
func (w *workServer) WorkServer() {
	fmt.Println("work server")
}

// Creating NewWorkServer
func NewWorkServer() (workServerArg *workServer) {
	return &workServer{}
}
