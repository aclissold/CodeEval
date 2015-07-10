/*
Problem 18: Multiples of a Number

Challenge Description:

Given numbers x and n, where n is a power of 2, print out the smallest
multiple of n which is greater than or equal to x. Do not use division
or modulo operator.

Input sample:

The first argument will be a text file containing a comma separated list
of two integers, one list per line. e.g.

13,8
17,16
Output sample:

Print to stdout, the smallest multiple of n which is greater than or equal
to x, one per line.
e.g.

16
32
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

func main() {
    lines := readLines()
    for _, line := range lines {
		// Retrieve x and n
		data := strings.Split(line, ",")
		x, err := strconv.Atoi(data[0])
		if err != nil {
			log.Fatal(err)
		}
		n, err := strconv.Atoi(data[1])
		if err != nil {
			log.Fatal(err)
		}
		// Initialize and increment the multiplier
		multiplier := 1
		for x > multiplier*n {
			if x > multiplier*n {
				multiplier++
			}
		}

		// Print the result
		fmt.Println(multiplier * n)
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
