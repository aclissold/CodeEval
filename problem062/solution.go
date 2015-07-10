/*
Problem 62: N Mod M

Challenge Description:

Given two integers N and M, calculate N Mod M (without using any inbuilt modulus operator).

Input sample:

Your program should accept as its first argument a path to a filename. Each line in this file
contains two comma separated positive integers. E.g.

	20,6
	2,3

You may assume M will never be zero.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// check(err) eliminates some boilerplate
func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	lines := readLines()
	for _, line := range lines {
		data := strings.Split(line, ",")
		// n is the dividend
		n, err := strconv.Atoi(data[0])
		check(err)
		// m is the divisor
		m, err := strconv.Atoi(data[1])
		check(err)
		// q is the quotient
		q := n / m
		// r is the remainder (more specifically, n % m)
		r := n - q*m
		// print the result
		fmt.Println(r)
	}
}

// readLines does the work bufio's scanner API would do, but CodeEval doesn't have Go 1.1 yet.
func readLines() (lines []string) {
	// Avoid a panic accessing os.Args[1]
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./problem062 <file>")
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
