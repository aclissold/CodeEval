/*
Problem 116: Morse Code

CHALLENGE DESCRIPTION:

You have received a text encoded with Morse code and want to decode it.

INPUT SAMPLE:

Your program should accept as its first argument a path to a filename. Input example is the following:

    .- ...- ..--- .-- .... .. . -.-. -..-  ....- .....
    -... .... ...--

Each letter is separated by space char, each word is separated by 2 space chars.

OUTPUT SAMPLE:

Print out decoded words. E.g.

    AV2WHIECX 45
    BH3

Your program has to support letters and digits only.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(morse(scanner.Text()))
	}
}

// Created by using jQuery selectors to console.log these mappings from Wikipedia, then further
// modified with Vim macros and gofmt.
var characters = map[string]string{
	".-":    "A",
	"-...":  "B",
	"-.-.":  "C",
	"-..":   "D",
	".":     "E",
	"..-.":  "F",
	"--.":   "G",
	"....":  "H",
	"..":    "I",
	".---":  "J",
	"-.-":   "K",
	".-..":  "L",
	"--":    "M",
	"-.":    "N",
	"---":   "O",
	".--.":  "P",
	"--.-":  "Q",
	".-.":   "R",
	"...":   "S",
	"-":     "T",
	"..-":   "U",
	"...-":  "V",
	".--":   "W",
	"-..-":  "X",
	"-.--":  "Y",
	"--..":  "Z",
	"-----": "0",
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",

	"*": " ", // special case
}

// morse takes a string of morse code characters and converts them to ISO basic Latin characters.
// Any occurences of "  " (double spaces) will be preserved as a single space in the result.
func morse(line string) string {
	latin := ""
	line = strings.Replace(line, "  ", " * ", -1)
	for _, code := range strings.Split(line, " ") {
		latin += characters[code]
	}
	return latin
}
