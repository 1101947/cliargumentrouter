package cmdline

import (
	"testing"
)

type testCase struct {
	originalFlags map[string]string
	posargs []string
	shouldPass bool
}

func flagsToPosargs(f map[string]string) []string {
	p := []string{}
	for k,v := range(f) {
		s := "--" + k 
		if v != "" {
			s = s + "=" + v
		}
		p = append(p, s)
	}
	return p
} 

func TestParse(t *testing.T) {
	goodFlags := map[string]string{"config": "~/.config/myapp.conf", "NoOptionFalg": ""}
	badFlags := map[string]string{"--": "", "--=": ""}
	goodCase := testCase{
		originalFlags: goodFlags,
		posargs: flagsToPosargs(goodFlags),
		shouldPass: true,
	}
	badCase := testCase{
		originalFlags: badFlags,
		posargs: flagsToPosargs(badFlags),
		shouldPass: false,
	}
	testCases := []testCase{goodCase, badCase}
	var root Arg
	for _,v := range(testCases) {
		args, err := Parse(v.posargs)
		if err != nil && v.shouldPass {
			t.Fatal("Parsing, got error")
		}
		if err == nil && !v.shouldPass {
			t.Fatal("Parsing, didn't get error")
		}
		if !v.shouldPass {
			continue
		}
		root = args[0]
		if len(root.Flags) != len(v.originalFlags) {
			t.Fatal("Different flags amount after parsing.")
		}
		for k,val := range(root.Flags) {
			orVal, ok := v.originalFlags[k]
			if !ok && v.shouldPass  {
				t.Fatal("Didn't find flag. Original flags: ", v.originalFlags, "parsed: ", root.Flags)
			}
			if val != orVal {
				t.Fatal("Different values.")
			}
		}
	}
}
