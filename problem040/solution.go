/*
Problem 40: Self Describing Numbers

Challenge Description:

A number is a self-describing number when (assuming digit positions are labeled 0 to N-1), the digit
in each position is equal to the number of times that that digit appears in the number.

Input sample:

The first argument is the pathname to a file which contains test data, one test case per line. Each line contains a positive integer. E.g.

	2020
	22
	1210

Output sample:

If the number is a self-describing number, print out 1. If not, print out 0. E.g.

	1
	0
	1

For the curious, here's how 2020 is a self-describing number: Position '0' has value 2 and there is
two 0 in the number. Position '1' has value 0 because there are not 1's in the number. Position '2'
has value 2 and there is two 2. And the position '3' has value 0 and there are zero 3's.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var m = map[string]int{
	"0": 0,
	"1": 0,
	"2": 0,
	"3": 0,
	"4": 0,
	"5": 0,
	"6": 0,
	"7": 0,
	"8": 0,
	"9": 0,
}

// process determines how many of each digit is found on the line and stores it in a map.
func process(line string) {
	// reset m
	for digit, _ := range m {
		m[digit] = 0
	}
	// count how many of each digit
	for _, ch := range line {
		m[string(ch)]++
	}
}

// check returns 1 if the line represents a self-describing number, and 0 if it doesn't.
func check(line string) int {
	for i, ch := range line {
		digit, err := strconv.Atoi(string(ch))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if m[(strconv.Itoa(i))] != digit {
			return 0
		}
	}
	// Did not find any incorrect values
	return 1
}

func main() {
	lines := readLines()
	for _, line := range lines {
		process(line)
		fmt.Println(check(line))
	}
}

// readLines does the work bufio's scanner API would do, but CodeEval doesn't have Go 1.1 yet.
func readLines() (lines []string) {
	// Avoid a panic accessing os.Args[1]
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./problem040 <file>")
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
