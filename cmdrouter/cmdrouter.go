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

type Router interface {
	Handle(path string, handler Handler) 
	HandleFunc(path string, fn func(posargs []string) error)
	Process(posargs []string) error
}
