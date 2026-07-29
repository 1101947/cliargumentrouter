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
		status: FlagsNotParsed(),
		aliasesToSetters: map[string]setter{},
		namesToFlagsPointers: map[string]*flag{},
	}
}

func (f *flags) Add(F FlagDescription) (*flag, error) {
	if len(F.Names) < 1 {
		return nil, fmt.Errorf("Flag have to have at least one name.")
	}
	for _,alias := range(F.Names) {
		if _, ok := f.aliasesToSetters[alias]; ok {
			fl, err := f.findFlagByAlias(alias)
			if err != nil {
				return nil, fmt.Errorf("Trying to find flag by alias, got: %w", err)
			}
			return nil, fmt.Errorf("Alias: '%s' is already in use by %s flag", alias, fl)
		} 
	}
	if _, ok := f.namesToFlagsPointers[F.Names[0]]; ok {
		return nil, fmt.Errorf("Name is already in use")
	}
	constructedFlag, err := constructFlag(F)
	if err != nil {
		return constructedFlag, fmt.Errorf("Constructing flag, got error: %w", err)
	} 
	f.namesToFlagsPointers[F.Names[0]] = constructedFlag 
	// TODO
	setter := setter{ flag: constructedFlag, }
	for _,alias := range(F.Names) {
		f.aliasesToSetters[alias] = setter
	}
	return constructedFlag, nil
}
type FlagDescription struct {
	Names []string
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

func (v valueSpecificationStatus) ValueMayBeSpecified() bool {
	if v == ValueSpecificationIsOptional() || v == ValueSpecificationIsRequired() {
		return true
	}
	return false
}

func (v valueSpecificationStatus) ValueShouldNotBeSpecified() bool {
	if v == ValueSpecificationIsForbiden() {
		return true
	}
	return false
}


func (f *flags) Parse(p []string) (int, error) {
	prefix := "--"
	keyValSep := '='
	var err error
	var key string
	var val string
	var valIsSpecified bool 
	var setter setter
	var ok bool
	var position int
	for position, arg := range(p) {
		if len(arg) < 3 {
			break
		}
		if arg[:2] != prefix {
			break
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
				}
				break
			}
		}
		if setter, ok = f.aliasesToSetters[key]; !ok {
			return position, fmt.Errorf("Got unknown flag: %s", key)
		}
		err = setter.Set(val, valIsSpecified)
		if err != nil {
			return position, fmt.Errorf("Setting flag, got: %s", err)
		}
	}
	f.status = Parsed()
	return position, nil 
}

//: isSet/isNotSet, valued/keyOnly/optional, assignedWithValue/NotAssignedWithValue

func (f flag) Get() (string, FlagStatus, error) {
	return f.value, f.getFlagStatus(), nil 
}

type FlagValue struct {
	ValueWasSet bool
	Value string
}

func (f flags) Status() flagsStatus {
	return f.status 
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
	status flagsStatus
	aliasesToSetters map[string]setter
	namesToFlagsPointers map[string]*flag
	//optional []flag
	//required []flag
}

func constructFlag(F FlagDescription) (*flag, error) {
	if len(F.Names) < 1 {
		return nil, fmt.Errorf("Flag must have at least one name.")
	}
	f := flag{
		name: F.Names[0],
		aliases: F.Names,
		value: F.DefaultValue,
		description: F.Description,
		isRequired: F.IsRequired,
		valueSpecificationStatus: F.ValueSpecificationStatus,
		isSet: false,
		isAssignedWithValue: false,
		isRead: false,
	}
	return &f, nil
} 

type setter struct {
	flag *flag
}

func (s setter) Set(v string, valueIsSpecified bool) error {
	if s.flag.isRead {
		return fmt.Errorf("Trying to set flag '%s' after it has been already read", s.flag.name)
	}
	if s.flag.isSet {
		return fmt.Errorf("Flag: %s has already been set.", s.flag.name)
	}
	s.flag.isSet = true
	if valueIsSpecified {
		s.flag.isAssignedWithValue = true
		if s.flag.valueSpecificationStatus.ValueShouldNotBeSpecified() {
			return fmt.Errorf("Value was specified to a flag that should have no value specified: %s.", s.flag.name)
		}
		s.flag.value = v
	}
	return nil
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

type FlagStatus struct {
	IsSet bool
	IsAssignedWithValue bool
	IsRead bool
} 

func (f flag) getFlagStatus() FlagStatus {
	return FlagStatus{
		IsSet: f.isSet,
		IsAssignedWithValue: f.isAssignedWithValue,
		IsRead: f.isRead,
	}

}


type flagsStatus string 
// notParsed, flagsParsed, flagsRead

