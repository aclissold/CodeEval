package main

import "testing"

func TestEvaluate(t *testing.T) {
	for input, want := range cases {
		got := evaluate(input.number, input.pattern)
		if got != want {
			t.Errorf("evaluate(%q, %q) == %d, want %d",
				input.number, input.pattern, got, want)
		}
	}
}

type input struct {
	number, pattern string
}

var cases = map[input]int{
	input{"3413289830", "a-bcdefghij"}: -413289827,
	input{"776", "a+bc"}:               83,
	input{"12345", "a+bcde"}:           2346,
	input{"1232", "ab+cd"}:             44,
	input{"90602", "a+bcde"}:           611,
}
