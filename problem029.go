/*

Problem 29: Unique Elements

Challenge Description:

You are given a sorted list of numbers with duplicates. Print out the sorted list with duplicates
removed.

Input sample:

File containing a list of sorted integers, comma delimited, one per line. E.g.

	1,1,1,2,2,3,3,4,4
	2,3,4,5,5

Output sample:

Print out the sorted list with duplicates removed, one per line.
E.g.

	1,2,3,4
	2,3,4,5

*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func contains(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

func main() {
	lines := readLines()
	for _, line := range lines {
		input := strings.Split(line, ",")
		var output = make([]string, 0, 10)
		for _, v := range input {
			if contains(output, v) {
				continue
			}
			output = append(output, v)
		}
		fmt.Println(strings.Join(output, ","))
	}
}

// readLines does the work bufio's scanner API would do, but CodeEval doesn't have Go 1.1 yet.
func readLines() (lines []string) {
	// Avoid a panic accessing os.Args[1]
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./problem029 <file>")
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
