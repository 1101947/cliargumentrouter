package flag

import (
	"testing"
)

type testCase struct {
	name string
	given
	expected
}

// No alarms and
func (t testCase) NoSurprises() bool {
	resultCmd, resultErr := t.given.dF.Parse(t.given.cmd)
	if resultErr != t.expected.err {
		return false
	}
	if !compareCmd(resultCmd, t.expected.cmd) {
		return false
	}
	if !(*t.given.dF).isEqualTo(*t.expected.dF) {
		return false
	}
	return true
} 

type given struct {
	cmd []string
	dF *defaultFlags
}

type expected struct {
	cmd []string
	dF *defaultFlags
	err error
}

func compareCmd(cmd1, cmd2 []string) bool {
	if len(cmd1) != len(cmd2) {
		return false
	}
	for i:=0;i<len(cmd1);i++ {
		if cmd1[i] != cmd2[i] {
			return false
		}
	}
	return true
}

func Test_DefaultFlagsParse(t *testing.T) {
	testCases := getTestCases()
	for _, v := range(testCases) {
		t.Run(v.name, func(t *testing.T) {
			if !v.NoSurprises() {
				t.Errorf("expected: %+v, \n got : %+v ", *v.expected.dF, (*(v.given.dF)))
			}
		})
	}
}

func getTestCases() []testCase {
	ddf := GetDefaultFlags()
	return  []testCase{
		testCase{
			name: "first",
			given: given{
				cmd: []string{"--flag1=7", "--flag3=value", "--flag2=value2", "--flag3=sd"},
				dF: &ddf,
			},
			expected: expected{
				cmd: []string{},
				dF: &defaultFlags{
					prefix: "--", nameValueSeparator: "=", kwargs: kwargs{
						"flag1": []pair{pair{value: "7", position: 0},},
						"flag3": []pair{pair{value: "value", position: 1},pair{value: "sd", position: 3},},

						"flag2": []pair{pair{value: "value2", position: 2},},
					},
				},
				err: nil,
			},

		},
	}
}
