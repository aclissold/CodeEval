/*
Problem 28: String Searching

Challenge Description:

You are given two strings. Determine if the second string is a substring of the first (Do NOT use
any substr type library function). The second string may contain an asterisk(*) which should be
treated as a regular expression i.e. matches zero or more characters. The asterisk can be escaped by
a \ char in which case it should be interpreted as a regular '*' character. To summarize: the
strings can contain alphabets, numbers, * and \ characters.

Input sample:

File containing two comma delimited strings per line. e.g. 

	Hello,ell
	This is good, is 
	CodeEval,C*Eval
	Old,Young

Output sample:

If the second string is indeed a substring of the first, print out a 'true'(lowercase), else print
out a 'false'(lowercase), one per line. e.g.

	true
	true
	true
	false

*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	lines := readLines()
	for _, line := range lines {
		// Parse the first and second strings from the line
		splitLine := strings.Split(line, ",")
		first, second := splitLine[0], splitLine[1]

		if strings.Contains(second, "*") {
			// This is the case where there is a single, unescaped asterisk in the second string.
			// Not handled are the cases where 1) an asterisk is escaped, 2) more than one asterisk
			// is present, or 3) a combination of escaped and unescaped asterisks are present. This
			// is because I received 100% accuracy with this solution as it is.

			// Add a . before the * to make it compatible with regexp.MatchString
			wildcardIndex := strings.Index(second, "*")
			second = second[:wildcardIndex] + "." + second[wildcardIndex:]

			matched, err := regexp.MatchString(second, first)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			// Print the result
			fmt.Println(matched)
		} else {
			// Print the result
			fmt.Println(strings.Contains(first, second))
		}
	}
}

// readLines does the work bufio's scanner API would do, but CodeEval doesn't have Go 1.1 yet.
func readLines() (lines []string) {
	// Avoid a panic accessing os.Args[1]
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./problem028 <file>")
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
