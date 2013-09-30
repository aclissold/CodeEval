/*
Problem 8: Reverse Words

Challenge Description:

Write a program to reverse the words of an input sentence.

Input sample:

The first argument will be a text file containing multiple sentences, one per line. Possibly empty
lines too. e.g.
    Hello World
    Hello CodeEval
    Output sample:
Print to stdout, each line with its words reversed, one per line. Empty lines in the input should be
ignored. Ensure that there are no trailing empty spaces on each line you print.  e.g.
    World Hello
    CodeEval Hello
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	lines := readLines()
	for _, line := range lines {
		// Create a list of words
		wordList := strings.Split(line, " ")
		var output string
		for _, word := range wordList {
			// Build a string in reversed-word order
			output = word + " " + output
		}
		// Remove the last character, always a space
		output = output[:len(output)-1]
		fmt.Println(output)
	}
}

// readLines does the work bufio's scanner API would do, but CodeEval doesn't have Go 1.1 yet.
func readLines() (lines []string) {
	// Avoid a panic accessing os.Args[1]
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./problem008 <file>")
		os.Exit(1)
	}

	// Read data from a file given as a command line argument
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Parse the lines of the file from data
	lines = strings.Split(string(data), "\n")

	// Remove the empty string after the final \n
	lines = lines[:len(lines)-1]

	return
}
