/*
Problem 131: Split the Number

CHALLENGE DESCRIPTION:

You are given a number N and a pattern. The pattern consists of lowercase latin
letters and one operation "+" or "-". The challenge is to split the number and
evaluate it according to this pattern e.g.  1232 ab+cd -> a:1, b:2, c:3, d:2 ->
12+32 -> 44

INPUT SAMPLE:

Your program should accept as its first argument a path to a filename. Each line
of the file is a test case, containing the number and the pattern separated by a
single whitespace. E.g.

	3413289830 a-bcdefghij
	776 a+bc
	12345 a+bcde
	1232 ab+cd
	90602 a+bcde

OUTPUT SAMPLE:

For each test case print out the result of pattern evaluation. E.g.

	-413289827
	83
	2346
	44
	611

Constraints:
N is in range [100, 1000000000]

*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		fmt.Println(evaluate(s[0], s[1]))
	}
}

// evaluate splits number at the index of + or - in pattern and performs that
// operation on the two halves.
func evaluate(number, pattern string) int {
	if i := strings.Index(pattern, "+"); i != -1 {
		l, r := splint(number, i)
		return l + r
	} else if i := strings.Index(pattern, "-"); i != -1 {
		l, r := splint(number, i)
		return l - r
	}

	log.Fatal("pattern must contain either + or -")
	return -1
}

// splint splits number at i into two ints (get it?).
func splint(number string, i int) (l, r int) {
	left, right := number[:i], number[i:]
	l, _ = strconv.Atoi(left)
	r, _ = strconv.Atoi(right)
	return l, r
}
