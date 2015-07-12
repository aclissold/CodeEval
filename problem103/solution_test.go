package main

import "testing"

func TestWinner(t *testing.T) {
	for s, want := range cases {
		got := winner(s)
		if got != want {
			t.Errorf("winner(%q) == %d, want %d", s, got, want)
		}
	}
}

var cases = map[string]int{
	"3 3 9 1 6 5 8 1 5 3":   5,
	"9 2 9 9 1 8 8 8 2 1 1": 0,
	"1":                 1,
	"9":                 1,
	"1 1":               0,
	"9 9":               0,
	"1 2 3 4 5 6 7 8 9": 1,
	"9 8 7 6 5 4 3 2 1": 9,
}
