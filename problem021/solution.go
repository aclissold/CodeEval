/*
Problem 21: Sum of Digits

Challenge Description:

Given a positive integer, find the sum of its constituent digits.

Input sample:

The first argument will be a text file containing positive integers, one per line. E.g.

    23
    496

Output sample:

Print to stdout, the sum of the numbers that make up the integer, one per line. E.g.

    5
    19

*/
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)

func main() {
	lines := readLines()
    // For line in file:
	for _, line := range lines {
        var sum int
        // For character in line...
        for i := range line {
            // ...append the integer representation of that character
            digit, err := strconv.Atoi((string(line[i])))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
            sum += digit
        }
        fmt.Println(sum)
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
