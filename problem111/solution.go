/*
Problem 111: Longest Word

CHALLENGE DESCRIPTION:

In this challenge you need to find the longest word in a sentence. If the
sentence has more than one word of the same length you should pick the first
one.

INPUT SAMPLE:

Your program should accept as its first argument a path to a filename. Input
example is the following

	some line with text
	another line

Each line has one or more words. Each word is separated by space char.

OUTPUT SAMPLE:

Print the longest word in the following way.

	some
	another

*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(longest(scanner.Text()))
	}
}

// longest returns the longest word in line.
func longest(line string) string {
	longest := ""
	for _, word := range strings.Split(line, " ") {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}
