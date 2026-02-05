package cliargumentrouter 

import (
	"strings"
	"fmt"
	"github.com/1101947/cliargumentrouter/cmdrouter"

)

type plainRouter map[string]cmdrouter.Handler

type plainRouterHandler struct{
	helpMsg string
}

func (d plainRouterHandler) Run(cmd []string) {
	fmt.Println(d.helpMsg)
}


func NewPlainRouter() plainRouter {
	router := plainRouter{}
	router[""] = plainRouterHandler{}
	return router

}

func (p plainRouter) Handle(path string, handler cmdrouter.Handler) {
	p[path] = handler
}

func (p plainRouter) findHandler(cmd string) cmdrouter.Handler {
	for path := cmd; path != ""; {
		if handler, ok := p[path]; ok {
			return handler
		}

		if i := strings.LastIndexByte(path, ' '); i > 0 {
			path = path[:i]
		} else {
			path = ""
		}
	}
	return p[""]
}

func (p plainRouter) Run(cmd []string) {
	path := strings.Join(cmd, " ")
	handler := p.findHandler(path)
	handler.Run(cmd)
}

func (p plainRouter) HandleFunc(path string, fn func(cmd []string)) {
	handler := cmdrouter.RunnerFunc(fn)
	p[path] = handler
}
