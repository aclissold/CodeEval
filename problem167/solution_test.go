package main

import "testing"

func TestTruncate(t *testing.T) {
	for line, want := range cases {
		got := truncate(line)
		if got != want {
			t.Errorf("truncate(%q) == %q, want %q", line, got, want)
		}
	}
}

var cases = map[string]string{
	"Tom exhibited.": "Tom exhibited.",
	"Amy Lawrence was proud and glad, and she tried to make Tom see it in her face - but he wouldn't look.": "Amy Lawrence was proud and glad, and... <Read More>",
	"Tom was tugging at a button-hole and looking sheepish.":                                                "Tom was tugging at a button-hole and looking sheepish.",
	"Two thousand verses is a great many - very, very great many.":                                          "Two thousand verses is a great many -... <Read More>",
	"Tom's mouth watered for the apple, but he stuck to his work.":                                          "Tom's mouth watered for the apple, but... <Read More>",
	"1234567891123456789212345678931234567894123456789512345":                                               "1234567891123456789212345678931234567894123456789512345",
	"12345678911234567892123456789312345678941234567895123456":                                              "1234567891123456789212345678931234567894... <Read More>",
	"123456789112345678921234567893123456789 1234567895123456":                                              "123456789112345678921234567893123456789... <Read More>",
	"12345678911234567892123456789312345678 41234567895123456":                                              "12345678911234567892123456789312345678... <Read More>",
}
