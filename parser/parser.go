package parser

import (
	//"github.com/1101947/cliargumentrouter/command"
)

type Parser interface {
	Serialize() 
	Deserialize() error
}

type DirtyParser interface {
	PurifyParser() (Parser, error)
}

//
//type CliParser interface {
//	Parser
//	GenerateCommand() command.Command
//}
