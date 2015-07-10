package main

import "testing"

func TestLongest(t *testing.T) {
	for line, want := range cases {
		got := longest(line)
		if got != want {
			t.Errorf("longest(%q) == %q, want %q", line, got, want)
		}
	}
}

var cases = map[string]string{
	"some line with text": "some",
	"another line":        "another",
}
