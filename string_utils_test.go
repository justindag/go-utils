package utils

import "testing"

func TestStrEmpty(t *testing.T) {
	testVals := []struct {
		s        string // input
		expected bool   // expected result
	}{
		{"", true},
		{"notempty", false},
		{" ", false},
	}

	for _, v := range testVals {
		actual := StrEmpty(v.s)
		if actual != v.expected {
			t.Errorf("StrEmpty(%s): expected '%v', actual '%v'", v.s, v.expected, actual)
		}
	}
}

func TestStrNotEmpty(t *testing.T) {
	testVals := []struct {
		s        string // input
		expected bool   // expected result
	}{
		{"", false},
		{"notempty", true},
		{" ", true},
	}

	for _, v := range testVals {
		actual := StrNotEmpty(v.s)
		if actual != v.expected {
			t.Errorf("StrNotEmpty(%s): expected '%v', actual '%v'", v.s, v.expected, actual)
		}
	}
}