package flag

import (
	"fmt"
)

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
	return flags{
		aliasesToSetters: map[string]setter{},
		namesToFlagsPointers: map[string]*flag{},
	}
}

func (f flags) Add(F FlagDescription) (*flag, error) {
	for _,alias := range(F.Aliases) {
		if alias == F.Name {
			return nil, fmt.Errorf("Name and alias matches: %s . Name and alias should not match", alias)
		}
		if _, ok := f.aliasesToSetters[alias]; ok {
			fl, err := f.findFlagByAlias(alias)
			if err != nil {
				return nil, fmt.Errorf("Trying to find flag by alias, got: %w", err)
			}
			return nil, fmt.Errorf("Alias: '%s' is already in use by %s flag", alias, fl)
		} 
	}
	if _, ok := f.namesToFlagsPointers[F.Name]; ok {
		return nil, fmt.Errorf("Name is already in use")
	}
	constructedFlag := constructFlag(F)
	f.namesToFlagsPointers[F.Name] = constructedFlag 
	// TODO
	setter := setter{ flag: constructedFlag, }
	for _,alias := range(F.Aliases) {
		f.aliasesToSetters[alias] = setter
	}
	return constructedFlag, nil
}

type setter struct {
	flag *flag
}

func (s setter) Set(v string) error {
	if s.flag.isSet {
		return fmt.Errorf("Flag: %s has already been set.", s.flag.name)
	}
	// TODO: set
	s.flag.value = v
	s.flag.isSet = true
	return nil
}

type FlagDescription struct {
	Name string
	Aliases []string
	IsRequired bool
	// value specification is required / value specification is forbiden / value specification is optional
	ValueSpecificationStatus valueSpecificationStatus 	
	Description string
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

func (f flags) Parse(p []string) {
}


type FlagStatus struct {
	IsSet bool
} 

//: isSet/isNotSet, valued/keyOnly/optional, assignedWithValue/NotAssignedWithValue

func (f flag) Get() (string, FlagStatus, error) {
	return "", FlagStatus{}, nil 
}

type FlagValue struct {
	ValueWasSet bool
	Value string
}

func (f flags) Status() flagsStatus {
	return FlagsNotParsed()
}

func (f flags) findFlagByAlias(a string) (string, error) {
	for name, val := range(f.namesToFlagsPointers) {
		for _, alias := range(val.aliases) {
			if a == alias {
				return name, nil
			}
		}
	}
	return "", fmt.Errorf("No flags have alias: %s", a)
}


type flags struct {
	aliasesToSetters map[string]setter
	namesToFlagsPointers map[string]*flag
	//optional []flag
	//required []flag
}

func constructFlag(F FlagDescription) *flag {
	f := flag{
		name: F.Name,
		aliases: F.Aliases,
		value: F.DefaultValue,
		description: F.Description,
		isRequired: F.IsRequired,
		valueSpecificationStatus: F.ValueSpecificationStatus,
		isSet: false,
		isAssignedWithValue: false,
		isRead: false,
	}
	return &f
} 

type flag struct {
	name string
	aliases []string
	value string
	description string
	valueSpecificationStatus valueSpecificationStatus
	isRequired bool
	isSet bool
	isAssignedWithValue bool
	isRead bool
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
