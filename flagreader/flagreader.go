package flagreader

import (
	"fmt"
	"errors"
)

func NoFlagWithThisNameWasFound() error {
	return errors.New("No flag with this name was found.")
}

func GetFlags(kwargs map[string]string) flags {
	f := flags{}
	var vc valueContainer
	for key, value := range(kwargs) {
		vc= valueContainer{
			value: value,
			haveBeenRead: false,
		}
		f[key] = vc 
	}
	return f
}

type flags map[string]valueContainer

type valueContainer struct {
	value string
	haveBeenRead bool
}

type Flags interface {
	AllFlagsHaveBeenRead() (bool, []string)
	ReadValueOf(string) (string, error)
}

func (f flags) AllFlagsHaveBeenRead() (bool, []string) {
	flagsNames := []string{}
	for name, value := range(f) {
		if !value.haveBeenRead {
			flagsNames = append(flagsNames, name)
		}
	}
	if len(flagsNames) > 0 {
		return false, flagsNames
	}
	return true, nil
}

func (f flags) ReadValueOf(name string) (string, error) {
	vc, ok := f[name]
	value := vc.value
	if !ok {
		return value, NoFlagWithThisNameWasFound() 
	}
	if vc.haveBeenRead {
		return value, fmt.Errorf("This flag have already been read.")
	}
	f[name] = valueContainer{
		haveBeenRead: true,
		value: value,
	}
	return value, nil 
}
