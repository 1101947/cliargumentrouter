package cmdrouter

import (
)

type Handler interface {
	Process(posargs []string) error
}

type ProcesserFunc func(posargs []string) error

func (R ProcesserFunc) Process(posargs []string) error {
	err := R(posargs)
	if err != nil {
		return err
	}
	return nil
}

//type DRouter map[string]Handler 
//
//
//func (D DRouter) Handle(posargs []string, H Handler) error {
//}
//
//func (D DRouter) HandleFunc(path []string, PF ProcesserFunc) error {
//}
//
//func (D DRouter) Process(posargs []string) error {
//	flags := flag.DefaultFlags("--", "=", posargs)
//	err := flags.Parse()
//	if err != nil {
//		log.Fatal(err)
//	}
//	_, posargs := flags.Extract() // TODO: do something with kwargs, mb process help
//	path := strings.Join(posargs, " ")
//	handler := d.findHandler(path)
//	handler.Run(posargs)
//}


