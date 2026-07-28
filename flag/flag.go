package flag

func NotRequired() string {
	return "Flag is not required"
}

func Parsed() flagsStatus {
	return flagsStatus("Flags are parsed")
}

func FlagsNotParsed() flagsStatus {
	return flagsStatus("Flags are not parsed")
}

func GetFlags() flags {
	return flags{}
}

func (f flags) Add(F Flag) error {
	return nil
}

type Flag struct {
	Name string
	Aliases []string
	IsRequired bool
	ValueSpecificationStatus valueSpecificationStatus // value specification is required / value specification is forbiden / value specification is optional
	DefaultValue string
}

type valueSpecificationStatus string

func ValueSpecificationIsRequired() valueSpecificationStatus {
	return valueSpecificationStatus("valueSpecificationIsRequired")
}

func ValueSpecificationIsForbiden() valueSpecificationStatus {
	return valueSpecificationStatus("valueSpecificationIsForbiden")
}

func ValueSpecificationIsOptional() valueSpecificationStatus {
	return valueSpecificationStatus("valueSpecificationIsOptional")
}


func (f flags) AddFlag(flagName, requirementStatus string, aliases []string) error {
	return nil
}

func (f flags) Parse(p []string) {
}


type FlagStatus struct {
	IsSet bool
} 

//: isSet/isNotSet, valued/keyOnly/optional, assignedWithValue/NotAssignedWithValue

func (f flags) Get(fName string) (string, FlagStatus, error) {
	return "", FlagStatus{}, nil 
}

type FlagValue struct {
	ValueWasSet bool
	Value string
}

func (f flags) Status() flagsStatus {
	return FlagsNotParsed()
}

type flags struct {
	//optional []flag
	//required []flag
}

type flagsStatus string 
// notParsed, flagsParsed, flagsRead


//
//// keyword arguments : {"flagName": "value", ... }
//type kwargs = map[string]kwarg
//type kwarg struct {
//	Aliases []string
//	Value string
//}
//type namedKwarg struct {
//	Name string
//	Value kwarg
//}
//
//// HERE STARTS
//type sometype struct {
//	f1 string
//	f2 string
//} 
//

//
//type flag struct {
//	name string
//	aliases []string
//	value string
//	isRequired bool
//	status
//}
//
//func (f flags) set(name, value string) error {
//	status := f.FlagStatus(name) 
//	if status == FlagDoesntExistStatus() {
//		return fmt.Errorf("Flag: %s doesn't exist.", name)
//	}
//	if status == FlagIsSet() {
//		return fmt.Errorf("Multiple %s flags encountered at positions: %d and %d. Multiple flags with the same name of the same command are not allowed", )
//	}
//}
//
//func (f flags) GetOptional(name string) (string, bool) {
//	flag := find(name)
//	if flag.IsSet() {
//		return flag.Value(), true
//	}
//	return "", false 
//}
//
//func (f flags) Parse(args []string) (int, error) {
//	flagPrefix := "--"
//	listOpener := "("
//	listCloser := ")"
//	keyValSep := "="
//	var err error
//	for position, arg := range(args) {
//		if arg[:len(flagPrefix)] == flagPrefix && arg[len(flagPrefix)+1] != listOpener {
//			key, value, err := parseFlag(arg[len(flagPrefix):])
//			f.add()
//			
//		} else if arg[0] == listOpener || arg[:len(flagPrefix)+len(listOpener)] == flagPrefix + listOpener {
//
//		} else if arg[0] != flagPrefix && arg[0] != listOpener {
//			return position, nil
//		}
//		if err != nil {
//			return position, fmt.Errorf("Iterating through positional arguments, got: %w", err)
//		}
//
//	}
//}
//
//func parseFlag(arg, keyValSep string) (key, value string, err error){
//	counter := 0
//	for i:=0;i<len(arg);i++ {
//		if arg[i] == keyValSep[counter] {
//			counter++
//		} else {
//			counter = 0
//		}
//		if counter == len(keyValSep) {
//			key = arg[:i-counter+1]
//			if len(key) == 0 {
//				return "", "", fmt.Errorf("Key must be specified") // TODO: add custom error type with separator symbol, and body fields
//			}
//			if len(key) < len(arg) {
//				value = arg[i+1:]
//			} else {
//				value = ""
//			}
//			return key, value, nil
//		}
//	}
//	return arg, "", nil
//}
//
//func (f flags) parsePosarg(arg string) 
//
//
//// To check if all flags are parsed and none fogotten before and after parsing posargs to flags and parsing flags to custom struct
//// befor parsing flags to custom type status must be: "posargs are parsed into flags"
//// and after parsing flags to custom type status must be: "flags are parsed into your type"
//func (f flags) GetStatus() flagsStatus {}
//
//func (f flags) GetRequired(name string) (string, error) {
//	flag := find(name)
//	if flag.IsSet() {
//		return flag.Value(), nil 
//	}
//	return "", ErrFlagNotFound() 
//}
//
//
//func (s sometype) ParseFlags(fl flags) {
//	f1StringVal, ok := fl.GetOptional("flagName")
//	if ok {
//		s.f1 = int(f1StringVal)
//	}
//	f2StringVal, err := fl.GetRequired("flagName")
//	if err != nil {
//		return err
//	}
//	s.f2 = int(f2StringVal)
//}
//// HERE ENDS 
//
//func isEqual(k1, k2 kwargs) bool {
//	if len(k1) != len(k2) {
//		return false
//	}
//	for key1, value1 := range(k1) {
//		value2, ok := k2[key1]
//		if !ok {
//			return false
//		}
//		if value1 != value2 {
//			return false
//		}
//		if len(mapOfFlagValues2) != len(mapOfFlagValues1) {
//			return false
//		}
//		for flagPosition1, flagValue1 := range mapOfFlagValues1 {
//			flagValue2, ok := mapOfFlagValues2[flagPosition1] 
//			if !ok {
//				return false
//			}
//			if flagValue1 != flagValue2 {
//				return false
//			}
//		} 
//	}
//	return true
//}
//
//// positional arguments
//type posargs = []string
//
//type Flags interface {
//	Parse(posargs) error
//	Extract() (kwargs, posargs) 
//}
