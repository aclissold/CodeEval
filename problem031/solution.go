/*
Problem 31: Rightmost Char

Challenge Description:

You are given a string 'S' and a character 't'. Print out the position of the rightmost occurrence
of 't' (case matters) in 'S' or -1 if there is none. The position to be printed out is zero based.

Input sample:

The first argument is a file, containing a string and a character, comma delimited, one per line.
Ignore all empty lines in the input file. E.g.

    Hello World,r
    Hello CodeEval,E
    Output sample:

Print out the zero based position of the character 't' in string 'S', one per line. Do NOT print out
empty lines between your output.  E.g.

    8
    10

*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	lines := readLines()
	for _, line := range lines {
		// Temporary variable for S and t
		St := strings.Split(line, ",")

        // S: the given string, t: the character to print the rightmost index of
		S, t := St[0], St[1]

        // Used in for loop
        var matchFound bool

        // Iterate over S in reverse since we're printing the rightmost char and want to stop as
        // soon as we find it
		for i := len(S) - 1; i > -1; i-- {
            // Ensure matchFound is false
            matchFound = false
			if string(S[i]) == t {
                // Found a match; print its index
                matchFound = true
				fmt.Print(i)
				break
			}
		}
        // Print -1 if no match was found
        if !matchFound {
            fmt.Print(-1)
        }
		fmt.Println()
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
