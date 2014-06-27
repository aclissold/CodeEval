/*
Problem 112: Swap Elements

CHALLENGE DESCRIPTION:

You are given a list of numbers which is supplemented with positions that have to be swapped.

INPUT SAMPLE:

Your program should accept as its first argument a path to a filename. Input example is the
following

	1 2 3 4 5 6 7 8 9 : 0-8
	1 2 3 4 5 6 7 8 9 10 : 0-1, 1-3

As you can see a colon separates numbers with positions.

Positions start with 0.

You have to process positions left to right. In the example above (2nd line) first you process 0-1,
then 1-3.

OUTPUT SAMPLE:

Print the lists in the following way.

	9 2 3 4 5 6 7 8 1
	2 4 3 1 5 6 7 8 9 10
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
		items := strings.Split(scanner.Text(), ":")
		s := strings.Split(strings.TrimSpace(items[0]), " ")
		allIndices := strings.Split(items[1], (","))
		for _, indices := range allIndices {
			i, j := parseIndices(indices)
			swap(s, i, j)
		}
		fmt.Println(strings.Trim(fmt.Sprint(s), "[]"))
	}
}

// parseIndices parses a string such as "1-3" into integer indices.
func parseIndices(indices string) (i, j int) {
	ij := strings.Split(strings.TrimSpace(indices), "-")
	i, _ = strconv.Atoi(ij[0])
	j, _ = strconv.Atoi(ij[1])
	return
}

// swap swaps the i- and j-th elements of s in-place.
func swap(s []string, i, j int) {
	s[i], s[j] = s[j], s[i]
}
