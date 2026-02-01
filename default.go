package cmdrouter

import (
	"strings"
	"fmt"
)


const (
	debug logLevel = "debug"
	info logLevel = "info"
	warn logLevel = "warn"
	erroR logLevel = "error"
)

type logLevel string

type defaultRouter map[string]Handler

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

func (d defaultRouter) Handle(path string, handler Handler) {
	d[path] = handler
}

func (d defaultRouter) findHandler(cmd string) Handler {
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
	path := strings.Join(cmd, " ")
	handler := d.findHandler(path)
	handler.Run(cmd)
}

func (d defaultRouter) HandleFunc(path string, fn func(cmd []string)) {
	handler := RunnerFunc(fn)
	d[path] = handler
}
