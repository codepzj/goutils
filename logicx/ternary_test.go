package logicx

import "testing"

func TestTernary(t *testing.T) {
	type testCase struct {
		isTrue bool
		expr1  any
		expr2  any
		output any
	}
	testCaseMap := map[string]testCase{
		"No1": {true, 1, 2, 1},
		"No2": {false, 1.1, 2.1, 2.1},
		"No3": {true, "a", "b", "a"},
	}
	for key, tc := range testCaseMap {
		t.Run(key, func(t *testing.T) {
			got := Ternary(tc.isTrue, tc.expr1, tc.expr2)
			if got != tc.output {
				t.Errorf("expected: %d,get: %d", tc.output, got)
			}
		})
	}
}
