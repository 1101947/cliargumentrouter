package flag

type Flags interface {
	Parse(cmd []string) error
	Extract() (map[string][]struct{value string; position int}, map[string]string) 
}
