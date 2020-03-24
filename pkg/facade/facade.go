package facade

type setServer interface {
	SetServer()
}

type workServer interface {
	WorkServer()
}

type Facade interface {
	Work()
}

type server struct {
	set  setServer
	work workServer
}

//Server functionality
func (s *server) Work() {
	s.set.SetServer()
	s.work.WorkServer()
}

// Server creating
func WorkingServer(setServer setServer, workServer workServer) Facade {
	return &server{
		set:  setServer,
		work: workServer,
	}
}
