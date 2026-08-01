package flagreader

import (
	"testing"

	"github.com/1101947/cliargumentrouter/cmdline"
)

func TestFlagsShouldPass(t *testing.T) {
	posargs := []string{"--config=~/.config/app", "action", "--someflag"}
	cmd, err := cmdline.Parse(posargs)
	if err != nil {
		t.Fatal("Error parsing cmdline")
	}
	rootFlags := cmd[0].Flags
	fls := GetFlags(rootFlags)
	value, err := fls.ReadValueOf("config")
	if err != nil {
		t.Fatal("Error reading flag value")
	}
	if value != "~/.config/app" {
		t.Fatal("Got value: ", value,  "but expected: ~/.config/app")
	}
	haveBeenRead, flagnamesThatHavent := fls.AllFlagsHaveBeenRead()
	if !haveBeenRead {
		t.Fatal("Not all flags have been read !: ", flagnamesThatHavent)
	}
} 

func TestFlagsShouldGetErrorNotAllFlagsHaveBeenRead(t *testing.T) {
	posargs := []string{"--config=~/.config/app", "--yetanotherflag=true", "action", "--someflag"}
	cmd, err := cmdline.Parse(posargs)
	if err != nil {
		t.Fatal("Error parsing cmdline")
	}
	rootFlags := cmd[0].Flags
	fls := GetFlags(rootFlags)
	value, err := fls.ReadValueOf("config")
	if err != nil {
		t.Fatal("Error reading flag value")
	}
	if value != "~/.config/app" {
		t.Fatal("Got value: ", value,  "but expected: ~/.config/app")

	}
	if fls.AllFlagsHaveBeenRead() {
		t.Fatal("All flags have been read !")
	}
}
