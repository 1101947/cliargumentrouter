package noname_flagcontroller

// flagdescriber

type FlagDescription struct {
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
		registeredNames: map[string]*FlagDescription{},
		registeredFlags: map[string]*FlagDescription{},
		flagS: fl,
	}
	return f
}

type flags struct {
	registeredNames map[string]*flag
	registeredFlags map[string]*flag
	flagS flagreader.flags
}

// TODO: should i take Flag or *Flag ?
func (f *flags) Register(F Flag) (string, error) {
	for _,name := range(F.Names) {
		if fl, ok := f.registeredNames[name]; ok {
			return fmt.Errorf("Name: %s have already been used by flag: %s", name, fl.Name[0])
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
	f.registeredFlags[F.Names[0]] = *fl
	for _, name := range(F.Names) {
		f.registeredNames[name] = *fl
	}
	return nil
}

func (f *flags) GetValueOf(name string) (string, error) {
	p, ok := f.registeredFlags[name]
	if !ok {
		return p.defaultValue, fmt.Errorf("Flag with this name wasn't registered.")
	}
	value := p.defaultValue 
	var err error
	counter := 0
	for _, name := range(p.names) {
		value, err = f.flagS.ReadValueOf(name)
		if !errors.Is(err, flagreader.NoFlagWithThisNameWasFound()) {
			continue
		}
		if err != nil {
			return value, fmt.Errorf("Searching for flag, got: %w", err)
		}
		counter++
	}
	if counter > 1 {
		return value, fmt.Errorf("Found several identical flags with different aliases.")
	}
	if counter == 0 {
		if p.IsRequired {
			return p.defaultValue, fmt.Errorf("Flag is required to be specified, but haven't found it.")
		}
	}
	return value, nil
}


func (f *flags) HaveReadAll() bool {
	for name,p := range(f.registeredFlags) {
		if p.isRequired && !p.haveBeenRead {
			return false
		}
	} 
	return f.flagS.AllFlagsHaveBeenRead()
}
