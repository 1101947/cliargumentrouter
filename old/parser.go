package cliargumentrouter

type ParserRouter struct {
	Actions map[string](parser.Parser)
	NextAction string 
}




type DefaultRootPreParser struct {
	posargs []string
	kwrgs kwargs
	Actions map[string](parser.Parser)
	NextAction string 
	Version string
	HelpCommandDescription string
}

type ExtendenedPreParser struct {
	DefaultRootPreParser
	LogLevel string
	Config string
	DryRun bool

}


func GetDefaultRootPreParser(kwrgs kwargs, posargs []string) DefaultRootPreParser {
	return DefaultRootPreParser{
		posargs: posargs,
		kwrgs: kwrgs,
		Actions 
	}
}

