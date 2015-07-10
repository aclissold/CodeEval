package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(w2d(scanner.Text()))
	}
}

var numerals = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// w2d converts a semicolon-separated list of English number words s to a non-separated list of
// Arabic numerals. For example, "one;two;three" returns "123".
func w2d(s string) string {
	words := strings.Split(s, ";")
	arabic := ""
	for _, word := range words {
		arabic += numerals[word]
	}
	return arabic
}
