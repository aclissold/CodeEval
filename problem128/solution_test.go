package main

import "testing"

func TestCompress(t *testing.T) {
	for sequence, want := range cases {
		got := compress(sequence)
		if got != want {
			t.Errorf("compress(%q) == %q, want %q", sequence, got, want)
		}
	}
}

var cases = map[string]string{
	"40 40 40 40": "4 40",
	"40 40 40 40 29 29 29 29 29 29 29 29 57 57 92 92 92 92 92 86 86 86 86 86 86 86 86 86 86": "4 40 8 29 2 57 5 92 10 86",
	"73 73 73 73 41 41 41 41 41 41 41 41 41 41":                                              "4 73 10 41",
	"1 1 3 3 3 2 2 2 2 14 14 14 11 11 11 2":                                                  "2 1 3 3 4 2 3 14 3 11 1 2",
	"7":              "1 7",
	"0 1 2 97 98 99": "1 0 1 1 1 2 1 97 1 98 1 99",
	"2 2 2 3 3 3":    "3 2 3 3",
}
