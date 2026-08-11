package flagcontroller 

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
	config := Flag{
		Names: []string{"config", "conf"},
		IsRequired: false,
		Description: "Flag that allows to specify path to the config file to load.",
		DefaultValue: "/etc/myapp/conf",
	}
	err = fls.Register(config)
	if err != nil {
		t.Fatal("Error registering flag.", err)
	}
	flagValue, err := fls.GetValueOf(config.Names[0])
	if err != nil {
		t.Fatal("Getting value of flag, expected value, got error: ", err)
	}
	if flagValue != "~/.config/app" {
		t.Fatal("Expected flag value to be equal to : ~/.config/app , but it is equal to: ", flagValue)
	}
	ok, notReadenFlags := fls.HaveReadAll()
	if !ok {
		t.Fatal("Expected to have all flags, readen, but havent read these: ", notReadenFlags)
	}
}
