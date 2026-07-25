package cmdrouter

import (
	"github.com/1101947/cliargumentrouter/cmd"
)

type Handler interface {
	Process(posargs []string) (cmd.Cmd, error)
}

type ProcesserFunc func(posargs []string) (cmd.Cmd, error)

func (R ProcesserFunc) Process(posargs []string) (cmd.Cmd, error) {
	cmd, err := R(posargs)
	if err != nil {
		return nil, err
	}
	return cmd, nil
}
