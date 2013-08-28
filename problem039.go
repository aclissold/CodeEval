/*
Problem 39: Happy Numbers

Challenge Description:

A happy number is defined by the following process. Starting with any positive integer, replace the
number by the sum of the squares of its digits, and repeat the process until the number equals 1
(where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for
which this process ends in 1 are happy numbers, while those that do not end in 1 are unhappy
numbers.

Input sample:

The first argument is the pathname to a file which contains test data, one test case per line. Each
line contains a positive integer. E.g.

	1 7 22

Output sample:

If the number is a happy number, print out 1. If not, print out 0. E.g

	1 1 0

For the curious, here's why 7 is a happy number: 7->49->97->130->10->1. Here's why 22 is NOT a
happy number: 22->8->64->52->29->85->89->145->42->20->4->16->37->58->89 ...
*/

package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

// foundIntegers is a string representing all computed numbers in a sequence
var foundIntegers string

func isHappy(integer string) int {
	if integer == "1" {
		// Base case: 1 is already happy
		return 1
	} else if strings.Contains(foundIntegers, integer) {
		// If any number has been checked before, we've entered an endless loop
		return 0
	} else {
		// Add the number to the string of checked numbers
		foundIntegers += integer + " " // Space-separated list

		// Compute the next number in the sequence and detect if that's happy instead
		var nextInt int
		for _, ch := range integer {
			// For each character in the string representing an integer
			digit, err := strconv.Atoi(string(ch))
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			// Compute the next number in the sequence as a rolling sum
			nextInt += int(math.Pow(float64(digit), 2))
		}
		return isHappy(strconv.Itoa(nextInt))
	}
}

func main() {
	lines := readLines()
	for _, integer := range lines {
		// Ensure foundIntegers is empty before starting
		foundIntegers = ""
		// Print the solution to the problem
		fmt.Println(isHappy(integer))
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
