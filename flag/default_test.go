package flag

import (
	"testing"
)

func compareDefaultFlagsStruct(f, s defaultFlags) bool {
	if f.prefix != s.prefix {
		return false
	}
	if f.nameValueSeparator != s.nameValueSeparator {
		return false
	}
	if len(f.flags) != len(s.flags) {
		return false
	}
	for k,vF := range(f.flags) {
		vS, ok := s.flags[k]
		if !ok {
			return false
		}
		if vS != vF {
			return false
		}
	}
	return true
}

func TEst_DefaultFlagsParse(t *testing.T) {
	type given struct{
		dF *defaultFlags
		cmd []string
	}
	type expected struct{
		dF defaultFlags
		cmd []string
		err error
	}
	type testCase struct{
		name string
		given
		expected
	}
	// TODO: add testcase where several identical flags present: progname --flag=value --flag=value --flag=value

	testCases := []testCase{
		{
			name: "default", 
			given: given{
				dF: &defaultFlags{
					prefix: "--", nameValueSeparator: "=", flags: map[string]string{},
				},
				cmd: []string{"--flag1", "--flag2=value2", "--flag3==sd"},
			},
			expected: expected{
				dF: defaultFlags{
					prefix: "--", nameValueSeparator: "=", flags: map[string]string{"flag1": "", "flag2": "value2", "flag3": "sd"},
				},
				cmd: []string{""},
				err: nil,
			},
		},
	}

	for _, v := range(testCases) {
		t.Run(v.name, func(t *testing.T) {
			rCmd, rErr := (v.given.dF).parse(v.given.cmd)
			for k,val := range(rCmd) {
				if val != v.expected.cmd[k] {
					t.Errorf("expected: %s, got: %s", v.expected.cmd[k], val)
				}
			}
			if rErr != v.expected.err {
				t.Errorf("expected: %v, got: %v", v.expected.err, rErr)
			}
			if compareDefaultFlagsStruct(*v.given.dF, v.expected.dF) {
				t.Errorf("expected %v, got: %v", v.expected.dF, *v.given.dF)
			}
		})
	}
}
