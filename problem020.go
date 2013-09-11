/*
Problem 20: Lowercase

Challenge Description:

Given a string write a program to convert it into lowercase.

Input sample:

The first argument will be a text file containing sentences, one per line. You can assume all characters are from the english language. e.g. 

HELLO CODEEVAL
This is some text
Output sample:

Print to stdout, the lowercase version of the sentence, each on a new line.
e.g.

hello codeeval
this is some text
*/
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
	for _, line := range lines() {
        fmt.Println(strings.ToLower(line))
    }
}

// lines does the work bufio's scanner API would do, but CodeEval doesn't have Go 1.1 yet.
// This is a minimal version of the function I normally use, for fun. It doesn't check if
// a file was actually passed on the command line (which would cause a panic accessing os.Args[1],
// nor does it check if an error is returned by ioutil.ReadFile.
func lines() (lines []string) {
	// Read data from a file given as a command line argument
	data, _ := ioutil.ReadFile(os.Args[1])

	// Parse the lines of the file from data
	lines = strings.Split(string(data), "\n")

	// Remove the empty string after the final \n
	lines = lines[:len(lines)-1]

	return
}
