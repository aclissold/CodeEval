/*
Problem 82: Armstrong Numbers

Challenge Description:

An Armstrong number is an n-digit number that is equal to the sum of the n'th powers of its digits.
Determine if the input numbers are Armstrong numbers.

Input sample:

Your program should accept as its first argument a path to a filename. Each line in this file has a
positive integer. E.g.

    6
    153
    351

Output sample:

Print out True/False if the number is an Armstrong number or not. E.g.

    True
    True
    False

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

// isArmstrong compares a given string representation of an integer to the sum of the nth powers
// of its digits and returns "True" if they are the same and false otherwise.
func isArmstrong(line string) string {
	// Create given and found to compare once found is computed
	given := line
	var found float64

	// n is used to compute the nth power of a digit
	n := len(given)

	// For each digit in the given number
	for _, digitRune := range given {
		// Convert the rune representing the digit to an actual number
		if digit, err := strconv.Atoi(string(digitRune)); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		} else {
			// Add the nth power of the digit to the rolling summation
			found += math.Pow(float64(digit), float64(n))
		}
	}

	// If given == found
	if given == strconv.Itoa(int(found)) {
		return "True"
	}
	return "False"
}

func main() {
	lines := readLines()
	for _, line := range lines {
		fmt.Println(isArmstrong(line))
	}
}

// readLines does the work bufio's scanner API would do, but CodeEval doesn't have Go 1.1 yet.
func readLines() (lines []string) {
	// Avoid a panic accessing os.Args[1]
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./problem082 <file>")
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
