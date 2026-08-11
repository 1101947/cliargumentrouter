package cmdrouter

import (
	"github.com/1101947/cliargumentrouter/executor"
)

type Handler interface {
	Process(posargs []string) (executor.Executor, error)
}

type ProcesserFunc func(posargs []string) (executor.Executor, error)

func (R ProcesserFunc) Process(posargs []string) (executor.Executor, error) {
	exec, err := R(posargs)
	if err != nil {
		return nil, err
	}
	return exec, nil
}
