package flag

// keyword arguments : "flagName": {flagGlobalPosition: "value", flagGlobalPosition: "value2"}
type kwargs = map[string]map[int]string

func isEqual(k1, k2 kwargs) bool {
	if len(k1) != len(k2) {
		return false
	}
	for flagKey1, mapOfFlagValues1 := range(k1) {
		mapOfFlagValues2, ok := k2[flagKey1]
		if !ok {
			return false
		}
		if len(mapOfFlagValues2) != len(mapOfFlagValues1) {
			return false
		}
		for flagPosition1, flagValue1 := range mapOfFlagValues1 {
			flagValue2, ok := mapOfFlagValues2[flagPosition1] 
			if !ok {
				return false
			}
			if flagValue1 != flagValue2 {
				return false
			}
		} 
	}
	return true
}

// positional arguments
type posargs = []string

type Flags interface {
	Parse(posargs) error
	Extract() (kwargs, posargs) 
}
