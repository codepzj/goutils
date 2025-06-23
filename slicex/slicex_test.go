package slicex

import (
	"slices"
	"testing"
)

func TestMap(t *testing.T) {
	type testCase struct {
		input  []int
		output []int
		fn     func(int) int
	}
	testCaseMap := map[string]testCase{
		"No1": {[]int{1, 2}, []int{3, 4}, func(num int) int { return num + 2 }},
		"No2": {[]int{1, 2}, []int{2, 4}, func(num int) int { return num * 2 }},
	}
	for key, tc := range testCaseMap {
		t.Run(key, func(t *testing.T) {
			got := Map(tc.input, tc.fn)
			if !slices.Equal(got, tc.output) {
				t.Errorf("expected: %d,get: %d", tc.output, got)
			}
		})
	}
}

func TestPush(t *testing.T) {
	type testCase struct {
		input  []int
		output []int
		num    int
	}
	testCaseMap := map[string]testCase{
		"No1": {[]int{1, 2}, []int{1, 2, 3}, 3},
	}
	for key, tc := range testCaseMap {
		t.Run(key, func(t *testing.T) {
			got := Push(tc.input, tc.num)
			if !slices.Equal(got, tc.output) {
				t.Errorf("expected: %d,get: %d", tc.output, got)
			}
		})
	}
}

func TestPop(t *testing.T) {
	type testCase struct {
		input   []int
		output  []int
		element int
	}
	testCaseMap := map[string]testCase{
		"No1": {[]int{1, 2}, []int{1}, 2},
		"No2": {[]int{3, 8}, []int{3}, 8},
	}
	for key, tc := range testCaseMap {
		t.Run(key, func(t *testing.T) {
			got, num := Pop(tc.input)
			if !slices.Equal(got, tc.output) {
				t.Errorf("expected: %d,get: %d", tc.output, got)
			}
			if num != tc.element {
				t.Errorf("expected pop number: %d,get number: %d", tc.element, num)
			}
		})
	}
}
