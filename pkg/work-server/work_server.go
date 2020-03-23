package work_server

type workServer struct {
	setServer bool
}

func (set *workServer) SetServer() bool {
	return true
}

func NewWorkServer() workServer {
	return workServer{
		setServer: true,
	}
}
