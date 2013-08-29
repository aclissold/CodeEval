/*
Problem 93: Capitalize Words
Capitalize Words  Share on LinkedIn

Challenge Description:

Write a program which capitalizes words in a sentence.

Input sample:

Your program should accept as its first argument a path to a filename. Input
example is the following

    Hello world
    javaScript language
    a letter

Output sample:

Print capitalized words in the following way.

    Hello World
    JavaScript Language
    A Letter
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
		newLine := ""
		for _, word := range strings.Split(line, " ") {
			newLine += strings.ToUpper(string(word[0])) + word[1:] + " "
		}
		newLine = newLine[:len(newLine)-1]
		fmt.Println(newLine)
	}
}

// readLines does the work bufio's scanner API would do, but CodeEval doesn't have Go 1.1 yet.
func readLines() (lines []string) {
	// Avoid a panic accessing os.Args[1]
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./XXX <file>") // XXX
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
