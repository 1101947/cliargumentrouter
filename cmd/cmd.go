package cmd

type Cmd interface {
	Exec() error
	Update()
}
