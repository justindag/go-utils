package utils

import (
	"testing"
)

func TestReverseStringSlice(t *testing.T) {
	testVals := []struct {
		in       []string // input
		expected []string // expected result
	}{
		{[]string{"1", "2", "3"}, []string{"3", "2", "1"}},
		{[]string{"2", "11", "3"}, []string{"3", "11", "2"}},
		{[]string{"2", "3", "1"}, []string{"1", "3", "2"}},
	}

	for _, v := range testVals {
		actual := ReverseStringSlice(v.in)
		assertSlice(t, actual, v.expected)
	}
}

func assertSlice(t *testing.T, actual, expected []string) {
	for i, s := range actual {
		if s != expected[i] {
			t.Errorf("ReverseStringSlice: expected '%v', actual '%v'", expected[i], s[i])
		}
	}
}
