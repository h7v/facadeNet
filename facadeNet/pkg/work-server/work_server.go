package work_server

import "fmt"

// WorkServer interface
type WorkServer interface {
	Work()
}

type workServer struct {
	workServer bool
}

// CheckWorkServer check how server is working
func (w *workServer) CheckWorkServer() (workServerArg string) {
	return "work server"
}

// WorkServer is how server working
func (w *workServer) WorkServer() {
	fmt.Println("work server")
}

// NewWorkServer creating 
func NewWorkServer() (workServerArg *workServer) {
	return &workServer{}
}
