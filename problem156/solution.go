/*
Problem 156: Roller Coaster

CHALLENGE DESCRIPTION:

You are given a piece of text. Your job is to write a program that sets the case
of text characters according to the following rules:

1. The first letter of the line should be in uppercase.
2. The next letter should be in lowercase.
3. The next letter should be in uppercase, and so on.
4. Any characters, except for the letters, are ignored during determination of letter case.

INPUT SAMPLE:

The first argument will be a path to a filename containing sentences, one per
line. You can assume that all characters are from the English language.

For example:

	To be, or not to be: that is the question.
	Whether 'tis nobler in the mind to suffer.
	The slings and arrows of outrageous fortune.
	Or to take arms against a sea of troubles.
	And by opposing end them, to die: to sleep.

OUTPUT SAMPLE:

Print to stdout the RoLlErCoAsTeR case version of the string.

For example:

	To Be, Or NoT tO bE: tHaT iS tHe QuEsTiOn.
	WhEtHeR 'tIs NoBlEr In ThE mInD tO sUfFeR.
	ThE sLiNgS aNd ArRoWs Of OuTrAgEoUs FoRtUnE.
	Or To TaKe ArMs AgAiNsT a SeA oF tRoUbLeS.
	AnD bY oPpOsInG eNd ThEm, To DiE: tO sLeEp.

CONSTRAINTS:

The length of each piece of text does not exceed 1000 characters.

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
		fmt.Println(rollercoaster(scanner.Text()))
	}
}

// rollercoaster converts s to RoLlErCoAsTeR cAsE.
func rollercoaster(s string) (converted string) {
	upper := true
	for _, r := range s {
		if r >= 'A' && r <= 'z' {
			if upper {
				converted += strings.ToUpper(string(r))
				upper = false
			} else {
				converted += strings.ToLower(string(r))
				upper = true
			}
		} else {
			converted += string(r)
		}
	}
	return converted
}
