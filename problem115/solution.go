/*
Problem 115: Mixed Content

CHALLENGE DESCRIPTION:

You have a string of words and digits divided by comma. Write a program which separates words with
digits. You shouldn't change the order elements.

INPUT SAMPLE:

Your program should accept as its first argument a path to a filename. Input example is the
following

    8,33,21,0,16,50,37,0,melon,7,apricot,peach,pineapple,17,21
    24,13,14,43,41

OUTPUT SAMPLE:

    melon,apricot,peach,pineapple|8,33,21,0,16,50,37,0,7,17,21
    24,13,14,43,41

As you can see you need to output the same input string if it has words only or digits only.
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
		fmt.Println(separate(scanner.Text()))
	}
}

// separate separates words from numbers in a comma-separated string
// and returns another comma-separated string with a pipe.
func separate(line string) string {
	words := make([]string, 0)
	numbers := make([]string, 0)

	for _, a := range strings.Split(line, ",") {
		if _, err := strconv.Atoi(a); err != nil {
			words = append(words, a)
		} else {
			numbers = append(numbers, a)
		}
	}

	wordString := strings.Join(words, ",")
	numberString := strings.Join(numbers, ",")

	if len(numbers) == 0 {
		return wordString
	} else if len(words) == 0 {
		return numberString
	}
	return wordString + "|" + numberString
}
