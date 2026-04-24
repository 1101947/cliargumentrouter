package cliargumentrouter 

import (
	"log"
	"fmt"
	"strings"
	"github.com/1101947/cliargumentrouter/cmdrouter"
	"github.com/1101947/cliargumentrouter/flag"
)


type NothingFoundHandler struct {}

func (N NothingFoundHanlder) Process(posargs []string) error {
	return nil
}

type parserRouter map[string]Handler

func NewParserRouter(defaultHandler Handler) parserRouter {
	router := parserRouter{}
	router[""] = defaultHandler 
	return router
}

func (p parserRouter) Handle(path string, handler Handler) {
	p[path] = handler
}

func (p parserRouter) findHandler(cmd string) cmdrouter.Handler {
	for path := cmd; path != ""; {
		if handler, ok := d[path]; ok {
			return handler
		}

		if i := strings.LastIndexByte(path, ' '); i > 0 {
			path = path[:i]
		} else {
			path = ""
		}
	}
	return d[""]
}

func (p parserRouter) Process(posargs []string) {
	path := strings.Join(posargs, " ")
	handler := d.findHandler(path)
	handler.Run(posargs)
}

func (d defaultRouter) HandleFunc(path string, fn func(posargs []string) error) {
	handler := RunnerFunc(fn)
	d[path] = handler
}
