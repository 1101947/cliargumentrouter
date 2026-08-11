package rootcmdrouter 

import (
	"github.com/1101947/cliargumentrouter/cmdrouter"
)

func NewRouter(posargs []string) RootCmdRouter {
	cmd, err := cmdline.Parse(posargs)
	if err != nil {
		return RootCmdRouter{}, fmt.Errorf("Parsing command line, got: %w", err)   
	}
	rootFlags := 
	return RootCmdRouter{
		sourceToParse: "",
		routes: defaultRoutes(),
		rootFlags: flagcontroller.GetFlags(cmd[0].Flags),
	}
}

func defaultRoutes() map[string]executor.Executor {
     return map[string]executor.Executor{"help": executor.Help()}
}

type RootCmdRouter struct {
	sourceToParse
	routes map[string]cmdrouter.Handler
	rootFlags 
}

func (R RootCmdRouter) Handle(name string, handler cmdrouter.Handler) error {
	if _, ok := R.routes[name]; ok {
		return fmt.Errorf("Handler with this name already exists.") 
	}
	R.routes[name] = handler
	return nil
} 

func (R RootCmdRouter) Process() (executor.Executor, error) {
} 
