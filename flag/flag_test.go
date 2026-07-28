package flag

import (
	"testing"
)

func TestFlags(t *testing.T) {
	//args := []string{"--flag", "--flag2=value", "--(flag3=val3 flag4)", "( flag5 flag6=val6 )", "posarg"}
	args := []string{"--flag", "--flag2=value"}
	flags := GetFlags()
	err := flags.AddFlag("flag", NotRequired(), []string{"flag", "fl"})
	if err != nil {
		t.Fatal("Adding flag: flag", "got: ", err)
	} 
	secondFlag := Flag{
		Name: "flag2",
		Aliases: []string{"flag2", "fl2"},
		IsRequired: false,
	}
	err = flags.Add(secondFlag)
	if err != nil {
		t.Fatal("Adding flag ", secondFlag, " got: ", err)
	}
	flags.Parse(args)
	sValFl1, status1, err := flags.Get("flag")
	if err != nil {
		t.Fatal("Getting flag with status: ", status1, " got: ", err)
	}
	sValFl2, status, err := flags.Get("flag2")
	if err != nil {
		t.Fatal("Getting flag with status: ", status, " got: ", err)
	}
	flagsStatus := flags.Status()
	if flagsStatus != Parsed() {
		t.Fatal("Flags weren't parsed: ", string(flagsStatus))
	}
	if !(status1.IsSet && sValFl2 == "value") {
		t.Fatal("Parsed flags wrong.:  ", sValFl1, sValFl2)
	}
} 
