package noname_flagcontroller

import (
		"github.com/1101947/cliargumentrouter/flagreader"
		"fmt"
		"errors"
)

// flagdescriber

type Flag struct {
	Names []string
	IsRequired bool
	Description string
	DefaultValue string
}

type flag struct {
	names []string
	isRequired bool
	haveBeenRead bool
	description string
	defaultValue string
}

func GetFlags(kwargs map[string]string) flags {
	fl := flagreader.GetFlags(kwargs)
	f := flags{
		registeredNames: map[string]*flag{},
		registeredFlags: map[string]*flag{},
		flagS: fl,
	}
	return f
}

type flags struct {
	registeredNames map[string]*flag
	registeredFlags map[string]*flag
	flagS flagreader.Flags
}

// TODO: should i take Flag or *Flag ?
func (f *flags) Register(F Flag) error {
	for _,name := range(F.Names) {
		if fl, ok := f.registeredNames[name]; ok {
			return fmt.Errorf("Name: %s have already been used by other flag:", fl.names[0])
		}
	}
	if _, ok := f.registeredFlags[F.Names[0]]; ok {
		return fmt.Errorf("This flag have alredy been registered.")
	}
	fl := flag{
		names: F.Names,  
		isRequired: F.IsRequired,
		haveBeenRead: false, 
		description: F.Description, 
		defaultValue: F.DefaultValue, 
	}
	f.registeredFlags[F.Names[0]] = &fl
	for _, name := range(F.Names) {
		f.registeredNames[name] = &fl
	}
	return nil
}

func (f *flags) GetValueOf(name string) (string, error) {
	p, ok := f.registeredFlags[name]
	if !ok {
		return p.defaultValue, fmt.Errorf("Flag with this name wasn't registered.")
	}
	var err error
	counter := 0
	value := p.defaultValue 
	iterValue := ""

	for _, name := range(p.names) {
		iterValue, err = f.flagS.ReadValueOf(name)
		if errors.Is(err, flagreader.NoFlagWithThisNameWasFound()) {
			err = nil
			continue
		}
//		if err != nil && err.Error() == flagreader.NoFlagWithThisNameWasFound().Error() {
//			err = nil
//			continue
//		}

		if err != nil {
			return value, fmt.Errorf("Searching for flag, got: %w", err)
		}
		counter++
		value = iterValue
	}
	if counter > 1 {
		return value, fmt.Errorf("Found several identical flags with different aliases.")
	}
	if counter == 0 {
		if p.isRequired {
			return p.defaultValue, fmt.Errorf("Flag is required to be specified, but haven't found it.")
		}
	}
	return value, nil
}


func (f flags) HaveReadAll() (bool, map[string]string) {
	notReadenFlags := map[string]string{}
	for name,p := range(f.registeredFlags) {
		if p.isRequired && !p.haveBeenRead {
			notReadenFlags[name] = "required, but not readen." 
		}
	} 
	ok, specifiedButNotReaden := f.flagS.AllFlagsHaveBeenRead()
	for _, flagname := range(specifiedButNotReaden) {
		notReadenFlags[flagname] = "specified, but not readen"
	}
	if len(notReadenFlags) == 0 && ok {
		return true, notReadenFlags 
	}
	return false, notReadenFlags 
}
