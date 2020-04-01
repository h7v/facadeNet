package facade

type setServer interface {
	SetServer()
}

type workServer interface {
	WorkServer()
}


// Facade interface
type Facade interface {
	Work()
}

type server struct {
	set  setServer
	work workServer
}

// Work implements server functionality
func (s *server) Work() {
	s.set.SetServer()
	s.work.WorkServer()
}

// WorkingServer is facade that implement server logic
func WorkingServer(setServer setServer, workServer workServer) (facadeArg Facade) {
	return &server{
		set:  setServer,
		work: workServer,
	}
}
