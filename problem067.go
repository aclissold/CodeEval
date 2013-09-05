/*
Problem 67: Hex to Decimal

Challenge Description:

You will be given a hexadecimal (base 16) number. Convert it into decimal (base 10).

Input sample:

Your program should accept as its first argument a path to a filename. Each line in this file contains a hex number. You may assume that the hex number does not have the leading 'Ox'. Also all alpha characters (a through f) in the input will be in lowercase. E.g.

	9f
	11

Output sample:

Print out the equivalent decimal number. E.g.

	159
	17

*/
package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

// hex is a mapping of base-16 characters to base-10 values
var hex = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"a": 10,
	"b": 11,
	"c": 12,
	"d": 13,
	"e": 14,
	"f": 15,
}

// reverse reverses a given string
func reverse(s string) (r string) {
	// Convert s to a slice of runes
	a := make([]rune, len(s))
	for i, ch := range s {
		a[i] = ch
	}

	// From http://golang.org/doc/effective_go.html#for
	// Reverse a
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	// Convert a back to a string
	for _, ch := range a {
		r += string(ch)
	}

	// Return reversed string
	return
}

func main() {
	// hexValue is the input
	var hexValue string
	// decimalValue is the output
	var decimalValue int

	lines := readLines()

	for _, line := range lines {
		decimalValue = 0
		// Assign hexValue in reverse so the for loop can use i for both the exponent and index
		hexValue = reverse(line)
		// Iterate over the (reversed) hex characters, building decimalValue with this formula:
		// decimalValue = sum of (hexCharacter * 16^position) for each character in hexValue
		for i := range hexValue {
			decimalValue += hex[string(hexValue[i])] * int(math.Pow(16, float64(i)))
		}
		// Print the results
		fmt.Println(decimalValue)
	}
}

// readLines does the work bufio's scanner API would do, but CodeEval doesn't have Go 1.1 yet.
func readLines() (lines []string) {
	// Avoid a panic accessing os.Args[1]
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./problem067 <file>")
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
