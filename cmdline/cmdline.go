package cmdline

type Posarg struct {
	Position int
	Name string
	Flags []Flag 
}

type Flag struct {
	Position int
	Key string
	IsValueSpecified bool
	Value string
}

// Everything that starts with -- is a flag
// Everything that doesn start with -- is a positional argument
// -- --= --KEY= are invalid flags
func Parse(args []string) ([]Posarg, error) {
	prefix := "--"
	keyValSep := '='
	invalidFlag := prefix + string(keyValSep)
	var err error
	var key string
	var val string
	var valIsSpecified bool 
	var ok bool
	var position int
	RootCmd := Posarg{
		Position: -1,
		Name: "root",
		Flags: []Flag{},
	}
	pArgs := []Posarg{RootCmd}
	var currParg Posarg
	var newFlag Flag
	for position, arg := range(args) {
		if len(arg) == 2 && arg == prefix {
			return pArgs, fmt.Errorf("Invalid syntax, got: %s", prefix)
		}
		if len(arg) == 3 && arg == invalidFlag {
			return pArgs, fmt.Errorf("Invalid syntax, got: %s", invalidFlag)
		}
		if len(arg) <= 2 || (len(arg) > 2 && arg[:2] != prefix) {
			currParg = Posarg{
				Position: position,
				Name: arg,
				Flags: []Flag{},
			}
			pArgs = append(pArgs, currParg)
			continue
		}
		key = arg[:2]
		val = ""
		valIsSpecified = false 
		for i,char := range(arg) {
			if char == keyValSep {
				key = arg[2:i]
				val = arg[i+1:]
				if len(val) > 0 {
					valIsSpecified = true
				} else {
					return pArgs, fmt.Errorf("Syntax error: Invalid flag syntax: --flag= . If flag specified with = symbol, value must be provided(inserted after =, like --flag=value.")
				}
				break
			}
		}
		newFlag = Flag{
			Position: position,
			Key: key
			IsValueSpecified: valueIsSpecified,
			Value: val,
		}
		(pArgs[len(pArgs)]).Flags = append(pArgs[len(pArgs)].Flags,  
	}
	return pArgs, nil 
}
