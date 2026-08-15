package cmdline

import (
	"fmt"
)

type Cmd struct {
	Name string
	Flags map[string]string
}
// Everything that starts with -- is a flag
// Everything that doesn start with -- is a positional argument
// -- --= --KEY= are invalid flags
// Takes cli arguments as array of strings and parses them starting with offset to command and its flags until new command is found, then returns parsed command and offset to cli arguments.
func Serialize(args []string, offset int) (Cmd, int, error) {
	prefix := "--"
	keyValSep := '='
	invalidFlag := prefix + string(keyValSep)
	var key string
	var val string
	cmd := Cmd{}
	cmd.Flags = map[string]string{}
	if offset < 0 {
		return cmd, offset, fmt.Errorf("Offset is a negative number. Offset must not be a negative number. Offset must be a positive number or zero.")
	} else if offset == 0 {
		cmd.Name = "root"
	} else {
		// TODO: offset vs offset + 1
		cmd.Name = args[offset] 
		offset++
	}
	var arg string
	for o:=offset;o<len(args);o++ {
		offset = o 
		arg = args[offset]
		if len(arg) == 2 && arg == prefix {
			return cmd, offset, fmt.Errorf("Invalid syntax, got: %s", prefix)
		}
		if len(arg) == 3 && arg == invalidFlag {
			return cmd, offset, fmt.Errorf("Invalid syntax, got: %s", invalidFlag)
		}
		if len(arg) <= 2 || (len(arg) > 2 && arg[:2] != prefix) {
			// TODO: or break ?
			return cmd, offset, nil 
		}
		key = arg[2:]
		val = ""
		for i,char := range(arg) {
			if char == keyValSep {
				key = arg[2:i]
				val = arg[i+1:]
				if len(val) <= 0 {
					return cmd, offset, fmt.Errorf("Syntax error: Invalid flag syntax: --flag= . If flag specified with = symbol, value must be provided(inserted after =, like --flag=value.")
				}
				break
			}
		}
		if _, ok := cmd.Flags[key]; ok {
			return cmd, offset, fmt.Errorf("This command already have flag with this key assigned: %s", key)
		}
		cmd.Flags[key] = val
	}
	return cmd, offset, nil
}
