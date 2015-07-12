package main

import (
	"testing"
)

func TestRollercoaster(t *testing.T) {
	for s, want := range cases {
		got := rollercoaster(s)
		if got != want {
			t.Errorf("rollercoaster(%q) == %q, want %q", s, got, want)
		}
	}
}

var cases = map[string]string{
	"To be, or not to be: that is the question.":   "To Be, Or NoT tO bE: tHaT iS tHe QuEsTiOn.",
	"Whether 'tis nobler in the mind to suffer.":   "WhEtHeR 'tIs NoBlEr In ThE mInD tO sUfFeR.",
	"The slings and arrows of outrageous fortune.": "ThE sLiNgS aNd ArRoWs Of OuTrAgEoUs FoRtUnE.",
	"Or to take arms against a sea of troubles.":   "Or To TaKe ArMs AgAiNsT a SeA oF tRoUbLeS.",
	"And by opposing end them, to die: to sleep.":  "AnD bY oPpOsInG eNd ThEm, To DiE: tO sLeEp.",
}
