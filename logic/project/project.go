package project

type Initializer interface {
	Init() error
}

func Init(i Initializer) error{
	return i.Init()
}


