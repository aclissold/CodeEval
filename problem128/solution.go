/*
Problem 128: Compressed Sequence

CHALLENGE DESCRIPTION:

Assume that someone dictates you a sequence of numbers, and you need to write it
down. For brevity, he dictates it as follows: first he dictates the number of
consecutive identical numbers, and then - the number itself.

For example:

The sequence below

	1 1 3 3 3 2 2 2 2 14 14 14 11 11 11 2

is dictated as follows:

	two times one, three times three, four times two, three times fourteen, three times eleven, one time two

and you have to write down the sequence

	2 1 3 3 4 2 3 14 3 11 1 2

Your task is to write a program that compresses a given sequence using this approach.

INPUT SAMPLE:

Your program should accept a path to a file as its first argument that contains
T number of lines. Each line is a test case represented by a sequence of
integers with the length L, where each integer is N separated by a space.

	40 40 40 40 29 29 29 29 29 29 29 29 57 57 92 92 92 92 92 86 86 86
	86 86 86 86 86 86 86
	73 73 73 73 41 41 41 41 41 41 41 41 41 41
	1 1 3 3 3 2 2 2 2 14 14 14 11 11 11 2
	7

OUTPUT SAMPLE:

For each test case, print out a compressed sequence of numbers separated by a
single space, one per line.

For example:

	4 40 8 29 2 57 5 92 10 86
	4 73 10 41
	2 1 3 3 4 2 3 14 3 11 1 2
	1 7

CONSTRAINTS:

* T is in a range from 20 to 50.
* N is in a range from 0 to 99.
* L is in a range from 1 to 400.

*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		fmt.Println(compress(scanner.Text()))
	}
}

// compress turns a sequence of numbers into a sequence of counts of numbers.
func compress(sequence string) (compressed string) {
	count := 0
	current := ""
	n := ""

	for _, n = range strings.Split(sequence, " ") {
		if current == n {
			count++
		} else if current == "" {
			current = n
			count++
		} else {
			compressed += fmt.Sprintf("%d %s ", count, current)
			current = n
			count = 1
		}
	}
	compressed += fmt.Sprintf("%d %s ", count, n)

	return compressed[:len(compressed)-1]
}
