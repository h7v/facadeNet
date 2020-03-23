package set_server

type setServer struct {
	setServer bool
}

func (set *setServer) SetServer() bool {
	return true
}

func NewSetServer() setServer {
	return setServer{
		setServer: true,
	}
}
