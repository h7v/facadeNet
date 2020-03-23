package facade

type setServer interface {
	setServer(bool) bool
}

type workServer interface {
	workServer(bool) bool
}

type facade struct {
	set  setServer
	work workServer
}

func (f *facade) Work(set bool, work bool) bool {
	if set && work {
		return true
	}
	return true
}

func Facade(setServer setServer, workServer workServer) *facade {
	return &facade{
		set:  setServer,
		work: workServer,
	}
}
