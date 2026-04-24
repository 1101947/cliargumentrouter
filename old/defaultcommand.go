package cliargumentrouter

import (
	"github.com/1101947/cliargumentrouter/command"
)
//
//type defaultCommand interface {
//	command.Command
//	registerFlag()
//	parseFlags(kwrgs kwargs)
//}

type kwargs = map[string]map[int]string

type DefaultCommand struct {
	HelpMsg string
	MainCommand func() error
}

func RunMain() error {
	err := MainCommand()
	return err
}
