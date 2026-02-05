package flag

type Flags interface {
	Parse(cmd []string) ([]string, error)
}

