package flag

import (
	"testing"
)

// TO GET: 
// ERRORS:
// flag is already added
// alias is already in use
// unknown flag
// required flag wasnt specified
// flags value wasnt specified
// flags value was specified
// not all flags were read
// required flags were not read
// non required flags were not read
func TestFlags(t *testing.T) {
	//args := []string{"--flag", "--flag2=value", "--(flag3=val3 flag4)", "( flag5 flag6=val6 )", "posarg"}
	args := []string{"--flag2=value"}
	flags := GetFlags()
	secondFlagDesc := FlagDescription{
		Name: "flag2",
		Aliases: []string{"fl2"},
		IsRequired: false,
		ValueSpecificationStatus: ValueSpecificationIsRequired(),
		Description: "just a flag",
		DefaultValue: "",
	}
	secondFlag, err := flags.Add(secondFlagDesc)
	if err != nil {
		t.Fatal("Adding flag ", secondFlagDesc, " got: ", err)
	}
	flags.Parse(args)
	sValFl2, status, err := secondFlag.Get()
	if err != nil {
		t.Fatal("Getting flag with status: ", status, " got: ", err)
	}
	flagsStatus := flags.Status()
	if flagsStatus != Parsed() {
		t.Fatal("Flags weren't parsed: ", string(flagsStatus))
	}
	if !(sValFl2 == "value") {
		t.Fatal("Parsed flags wrong.:  ", sValFl2)
	}
} 
