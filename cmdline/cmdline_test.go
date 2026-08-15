package cmdline

import (
	"testing"
)

func TestCmdline(t *testing.T) {
	line := []string{"--flag1", "--flag2=2"}
	cmd, offset, err := Serialize(line, 0)
	if offset <= 0 {
		t.Fatal("Expected offset to greater than 0, but is lesser or equal to zero: offset: ", offset, " Error: ", err)
	}
	if err != nil {
		t.Fatal(err)
	}

	if cmd.Name != "root" {
		t.Fatal("Expected command name to be root")
	}
	firstFlag, ok := cmd.Flags["flag1"]
	if !ok {
		t.Fatal("Expected to get flag1 value specified, but didnt")
	}
	if firstFlag != "" {
		t.Fatal("Expected first flag value to be \"\" , but is: ", firstFlag)
	}
	secondFlag, ok := cmd.Flags["flag2"]
	if !ok {
		t.Fatal("Expected to get flag1 value specified, but didnt")
	}
	if secondFlag != "2" {
		t.Fatal("Expected second flag value to be 2, but is: ", secondFlag)
	}
}

