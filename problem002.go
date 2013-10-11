/*
Problem 2: Longest Lines (see golang.org/pkg/sort/)

Challenge Description:

Write a program to read a multiple line text file and write the 'N' longest lines to stdout. Where
the file to be read is specified on the command line.

Input sample:

Your program should read an input file (the first argument to your program). The first line contains
the value of the number 'N' followed by multiple lines. You may assume that the input file is
formatted correctly and the number on the first line i.e. 'N' is a valid positive integer.e.g.

	2
	Hello World

	CodeEval
	Quick Fox
	A
	San Francisco

NOTE: For solutions in JavaScript, assume that there are 8 lines of input (i.e. line 1 will be N and
the next 7 lines will be the input lines

Output sample:

The 'N' longest lines, newline delimited. Do NOT print out empty lines. Ignore all empty lines in
the input. Ensure that there are no trailing empty spaces on each line you print. Also ensure that
the lines are printed out in decreasing order of length i.e. the output should be sorted based on
their length e.g.

	San Francisco
	Hello World

*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// line represents an individual line of a file.
type line struct {
	content string
	length  int
}

// lines contains all lines of a file.
var lines = make([]line, 0)

// by is the type of a "less" function that defines the ordering of its line arguments.
type by func(l1, l2 *line) bool

func (by by) sort(lines []line) {
	ls := &lineSorter{
		lines: lines,
		by:    by, // The sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ls)
}

// lineSorter joins a by function and a slice of lines to be sorted.
type lineSorter struct {
	lines []line
	by    func(l1, l2 *line) bool // Closure used in the less method.
}

// Len is part of sort.Interface.
func (s *lineSorter) Len() int {
	return len(s.lines)
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *lineSorter) Less(i, j int) bool {
	return s.by(&s.lines[i], &s.lines[j])
}

// Swap is part of sort.Interface.
func (s *lineSorter) Swap(i, j int) {
	s.lines[i], s.lines[j] = s.lines[j], s.lines[i]
}

func main() {

	// Closure that orders the line structure
	length := func(l1, l2 *line) bool {
		return l1.length > l2.length
	}

	// data are the unprocessed contents of a file.
	data := readLines()
	n, err := strconv.Atoi(data[0]); if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < len(data); i++ {
		nextLine := line{data[i], len(data[i])}
		lines = append(lines, nextLine)
	}

	by(length).sort(lines)

	for i := 0; i < n; i++ {
		fmt.Println(lines[i].content)
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
