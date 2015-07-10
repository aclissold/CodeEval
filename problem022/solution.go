/*
Problem 22: Fibonacci Series

Challenge Description:

The Fibonacci series is defined as:

    F(0) = 0; F(1) = 1; F(n) = F(n-1) + F(n-2) when n>1;.

Given a positive integer 'n', print out the F(n).

Input sample:

The first argument will be a text file containing a positive integer, one per line. e.g.

    5
    12

Output sample:

Print to stdout, the fibonacci number, F(n).
e.g.

    5
    144
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
		// n = int(line)
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
            // Base case
			fmt.Println(0)
		} else if n == 1 {
            // Base case
			fmt.Println(1)
		} else {
            // Fib
			a, b := 0, 1
			for i := 1; i < n; i++ {
				a, b = b, a+b
			}
			fmt.Println(b)
		}
	}
}

// readLines does the work bufio's scanner API would do, but CodeEval doesn't have Go 1.1 yet.
func readLines() (lines []string) {
	// Avoid a panic accessing os.Args[1]
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./problem020 <file>")
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
