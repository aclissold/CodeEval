/*
Problem 92: Penultimate Word

CHALLENGE DESCRIPTION:

Write a program which finds the next-to-last word in a string.

INPUT SAMPLE:

Your program should accept as its first argument a path to a filename. Input example is the
following

    some line with text
    another line
    Each line has more than one word.

OUTPUT SAMPLE:

Print the next-to-last word in the following way.

	with
	another
*/

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
