package flag

type flags interface {
	parse(cmd []string) []string 
}

