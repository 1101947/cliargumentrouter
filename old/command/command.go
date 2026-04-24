package command 

import ()

type Command interface {
	GetRunner() (func() error, error)
}

type DirtyCommand interface {
	PurifyCommand() (Command, error)
}
//
//type Commands interface {
//	Merge()
//	ExtractMergeResult() Command
//}
//
//type Parser interface {
//	Serialize() []byte
//	Deserialize() (Command, error)
//}
//
//type RunnerFunc func() error
//
//type SuperCommand interface {
//	Command
//	Handle(name string, cmd Command)
//	SetNextToExecute(name string) error
//}
