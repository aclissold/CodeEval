/*

Problem 19: Bit Positions

Challenge Description:

Given a number n and two integers p1,p2 determine if the bits in position p1 and p2 are the same or
not. Positions p1 and p2 and 1 based.

Input sample:

The first argument will be a text file containing a comma separated list of 3 integers, one list per
line. E.g.

	86,2,3
	125,1,2

Output sample:

Print to stdout, 'true'(lowercase) if the bits are the same, else 'false'(lowercase). E.g.

	true
	false

*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// check will cause the program to fail on non-nil errors.
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	lines := readLines()
	// For each line in the file
	for _, line := range lines {
		// Parse the line into a slice whose indices correspond to n, p1, and p2, respectively
		values := strings.Split(line, ",")
		n, err := strconv.Atoi(values[0]); check(err)
		p1, err := strconv.Atoi(values[1]); check(err)
		p2, err := strconv.Atoi(values[2]); check(err)

		// Store the larger of p1 and p2 in "max"
		var max int
		if p1 < p2 {
			max = p2
		} else {
			max = p1
		}

		var r, r1, r2 int
		// For each bit position in the binary representation of n up to position "max"
		for i := 1; i < max; i++ {
			// Store the remainder of n % 2 (which is the bit at position i)
			r = n % 2
			if i == p1 {
				// Store position 1's bit into r1
				r1 = r
			} else if i == p2 {
				// Store position 2's bit into r2
				r2 = r
			}
		}

		// Print the results
		fmt.Println(r1 == r2)
	}
}

// readLines does the work bufio's scanner API would do, but CodeEval doesn't have Go 1.1 yet.
func readLines() (lines []string) {
	// Avoid a panic accessing os.Args[1]
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./problem019 <file>")
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
