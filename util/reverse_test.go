package util

import "testing"

type testPair struct {
	in, out string
}

var tests = []testPair{
	{ "", "" },
	{ "x", "x" }, 
	{ "abcd", "dcba"},
}

func TestReverse(t *testing.T) {
	for _, tp := range tests {
		r := Reverse(tp.in)
		if r != tp.out {
			t.Errorf("called Reverse(%q), expected %q, got %q", tp.in, tp.out, r);
		}
	}
}
