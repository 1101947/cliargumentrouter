package cliargumentrouter

import (
	"github.com/1101947/cliargumentrouter/command"
)


type SuperCommand struct {
	name string
	defaultCommand command.Command
	commands []*SuperCommand
	next *SuperCommand
}

func (C Command) GetRunner() (func() error, error) {
	next, err := findNext(C.commands, C.next)
	if err != nil {
		return nil, fmt.Errorf("%s, got: %w", C.name, err) 
	}
	if next != nil {
		runner, err := &next.GetRunner()
		if err != nil {
			return nil, fmt.Errorf("%s, got: %w", C.name, err)
		}
		return runner, nil
	}
	runner, err := C.defaultCommand.GetRunner()
	if err != nil {
		return nil, fmt.Errorf("%s, got: %w", C.name, err)
	}
	return runner, nil
}

func findNext(commands []*Command, next *Command) (*Command, error) {
	counter := 0
	var nextC *Command
	for k,v := range commands {
		if v == next {
			counter++
			nextC = next
		}
	}
	if counter > 1 {
		// TODO: make more understandable
		return nil, fmt.Errorf("Command dublicate in commands pointers list")
	} else if counter = 1 {
		return nextC, nil
	} else {
		return nil, nil
	}
}

type Parser struct {
	data []byte
	parsers []*Parser
	next *Parser
	command *Command
}

func (P Parser) Serialize() []byte {
}

func (P Parser) Deserialize() (Command, error) {
	cmd, leftovers, err := parsible.Deserialize()
	P.command = cmd



}
