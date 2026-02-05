package cliargumentrouter 

import (
	"log"
	"fmt"
	"strings"
	"github.com/1101947/cliargumentrouter/cmdrouter"
	"github.com/1101947/cliargumentrouter/flag"
)

const (
	debug logLevel = "debug"
	info logLevel = "info"
	warn logLevel = "warn"
	erroR logLevel = "error"
)

type logLevel string

type defaultRouter map[string]cmdrouter.Handler

type defaultHandler struct{
	helpMsg string
	logLevel
	dryRun bool
}

func (d defaultHandler) Run(cmd []string) {
	fmt.Println(d.helpMsg)
}


func NewDefaultRouter() defaultRouter {
	router := defaultRouter{}
	router[""] = defaultHandler{}
	return router

}

func (d defaultRouter) Handle(path string, handler cmdrouter.Handler) {
	d[path] = handler
}

func (d defaultRouter) findHandler(cmd string) cmdrouter.Handler {
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

func (d defaultRouter) Run(cmd []string) {
	flags := flag.DefaultFlags("--", "=", cmd)
	err := flags.Parse()
	if err != nil {
		log.Fatal(err)
	}
	_, posargs := flags.Extract() // TODO: do something with kwargs, mb process help
	path := strings.Join(posargs, " ")
	handler := d.findHandler(path)
	handler.Run(posargs)
}

func (d defaultRouter) HandleFunc(path string, fn func(cmd []string)) {
	handler := cmdrouter.RunnerFunc(fn)
	d[path] = handler
}
