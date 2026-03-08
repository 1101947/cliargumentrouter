package cliargumentrouter

import ()

type Command interface {
	Run(kwrgs kwargs, posargs []string) error
}

type kwargs map[string]map[int]string

//type defaultCommand struct {}
type DefaultCommand interface {
	GetHelp() string
	RegisterPosargsProcessor(P PosargsProcessor) error
	RegisterFlag(F Flag) error
	Command
}

type Flag struct {
	Name string
	Required bool
	HandleMultiple OnMultiple	
}

type OnMultiple string

const (
	OnMultipleFail OnMultiple = "fail"
	OnMultipleOverwrite OnMultiple = "overwrite"
)

type PosargsProcessor interface {
	Run(posargs []string) error
}


