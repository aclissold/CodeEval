/*
Problem 147: Lettercase Percentage Ratio

CHALLENGE DESCRIPTION:

Your goal is to find the percentage ratio of lowercase and uppercase letters in line below.

INPUT SAMPLE:

Your program should accept as its first argument a path to a filename. Each line of input contains a
string with uppercase and lowercase letters E.g.:

    thisTHIS
    AAbbCCDDEE
    N
    UkJ

OUTPUT SAMPLE:

For each line from input, print the percentage ratio of uppercase and lowercase letters rounded to
the second digit after the dot. E.g.:

    lowercase: 50.00 uppercase: 50.00
    lowercase: 20.00 uppercase: 80.00
    lowercase: 0.00 uppercase: 100.00
    lowercase: 33.33 uppercase: 66.67
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lower, upper := percentages(scanner.Text())
		fmt.Printf("lowercase: %.2f uppercase: %.2f\n", lower, upper)
	}
}

const a = 97 // lowercase a's ASCII code

// percentages returns the % lowercase and uppercase letters in a given string
func percentages(line string) (lower float64, upper float64) {
	lowerCount, upperCount := 0.0, 0.0
	total := 0.0
	for _, char := range line {
		if char < a {
			upperCount++
		} else {
			lowerCount++
		}
		total++
	}
	return 100 * (lowerCount / total), 100 * (upperCount / total)
}
