package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])
	// Assume args are passed properly
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(lastWord(scanner.Text()))
	}
}

// lastWord returns the last word of a space-separated string of words.
// Assumes each line has more than one word.
func lastWord(s string) string {
	words := strings.Split(s, " ")
	return words[len(words)-2]
}
