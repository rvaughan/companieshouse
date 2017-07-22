package stringtools_test

import (
	"testing"
	"github.com/BalkanTech/companieshouse/api/stringtools"
)

func TestTitledString_String(t *testing.T) {
	var tests = []struct {
		s string
		expected string
	}{
		{"test", "Test"},
		{"this-is-a-test", "This is a test"},
	}

	for i, test := range tests {
		r := stringtools.TitledString(test.s)
		if  r.String() != test.expected {
			t.Fatalf("%d: %s Expected %q, but got %q.", i, test.s, test.expected, r)
		}
		t.Logf("%d: %s Expected %q.", i, test.s, test.expected)
	}
}

func TestReplaceBetween(t *testing.T) {
	os := "This is a **test** string"
	r := os
	e := "This is a [test] string"
	stringtools.ReplaceBetween(&r, "**", "[", "]")
	if  r != e {
		t.Fatalf("%s: Expected %q, but got %q.", os, e, r)
	}
	t.Logf("%s: Expected %q.", os, e)
}