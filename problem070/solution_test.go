package main

import "testing"

func TestOverlaps(t *testing.T) {
	for coordinates, want := range cases {
		got := overlaps(coordinates)
		if got != want {
			t.Errorf("overlaps(%q) == %q, want %q", coordinates, got, want)
		}
	}
}

var cases = map[string]string{
	"-3,3,-1,1,1,-1,3,-3": "False",
	"-3,3,-1,1,-2,4,2,2":  "True",
	"0,0,1,1,1,0,2,1":     "True",
}
