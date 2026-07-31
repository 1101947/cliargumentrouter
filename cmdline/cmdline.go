package cmdline

import (
	"fmt"
)

type Arg struct {
	Name string
	Flags map[string]string
}
// Everything that starts with -- is a flag
// Everything that doesn start with -- is a positional argument
// -- --= --KEY= are invalid flags
func Parse(args []string) ([]Arg, error) {
	prefix := "--"
	keyValSep := '='
	invalidFlag := prefix + string(keyValSep)
	var key string
	var val string
	RootCmd := Arg{
		Name: "root",
		Flags: map[string]string{},
	}
	pArgs := []Arg{RootCmd}
	var currParg Arg
	for _, arg := range(args) {
		if len(arg) == 2 && arg == prefix {
			return pArgs, fmt.Errorf("Invalid syntax, got: %s", prefix)
		}
		if len(arg) == 3 && arg == invalidFlag {
			return pArgs, fmt.Errorf("Invalid syntax, got: %s", invalidFlag)
		}
		if len(arg) <= 2 || (len(arg) > 2 && arg[:2] != prefix) {

			currParg = Arg{
				Name: arg,
				Flags: map[string]string{},
			}
			pArgs = append(pArgs, currParg)
			continue
		}
		key = arg[2:]
		val = ""
		for i,char := range(arg) {
			if char == keyValSep {
				key = arg[2:i]
				val = arg[i+1:]
				if len(val) <= 0 {
					return pArgs, fmt.Errorf("Syntax error: Invalid flag syntax: --flag= . If flag specified with = symbol, value must be provided(inserted after =, like --flag=value.")
				}
				break
			}
		}
		if _, ok := pArgs[len(pArgs)-1].Flags[key]; ok {
			return pArgs, fmt.Errorf("This command already have flag with this key assigned: %s", key)
		}
		(pArgs[len(pArgs)-1]).Flags[key] = val  
	}
	return pArgs, nil 
}
