/*

Problem 19: Bit Positions

Challenge Description:

Given a number n and two integers p1,p2 determine if the bits in position p1 and p2 are the same or
not. Positions p1 and p2 are 1 based.

Input sample:

The first argument will be a path to a filename containing a comma separated list of 3 integers, one
list per line. E.g.

	86,2,3
	125,1,2

Output sample:

Print to stdout, 'true'(lowercase) if the bits are the same, else 'false'(lowercase). E.g.

	true
	false

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), ",")
		n, _ := strconv.Atoi(tokens[0])
		p1, _ := strconv.Atoi(tokens[1])
		p2, _ := strconv.Atoi(tokens[2])
		fmt.Println(samebits(n, p1, p2))
	}
}

// samebits returns whether n has the same bits at positions p1 and p2 (1-based).
func samebits(n, p1, p2 int) bool {
	s := strconv.FormatInt(int64(n), 2)
	return s[len(s)-p1] == s[len(s)-p2]
}
