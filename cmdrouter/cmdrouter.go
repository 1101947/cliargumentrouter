package cmdrouter

import (
	//"strings"
	//"fmt"
)

type Handler interface {
	Run(cmd []string)
}

type RunnerFunc func(cmd []string)

func (R RunnerFunc) Run(cmd []string) {
	R(cmd)
}

type Router interface {
	Handle(path string, handler Handler) 
	HandleFunc(path string, fn func(cmd []string))
	Run(cmd []string)
}
