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
