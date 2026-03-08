package parser

import (
	"github.com/1101947/cliargumentrouter/command"
)

type Parser interface {
	Serialize() 
	Deserialize()
}

type CliParser interface {
	Parser
	GenerateCommand() command.Command
}
