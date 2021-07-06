package work_server

import "fmt"

type WorkServer interface {
	Work()
}

type workServer struct {
	workServer bool
}

//Work of the server
func (w *workServer) WorkServer() {
	fmt.Println("work server")
}

// Creating NewWorkServer
func NewWorkServer() *workServer {
	return &workServer{}
}
