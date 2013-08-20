/*
Problem 8: Reverse Words

Challenge Description:

Write a program to reverse the words of an input sentence.

Input sample:

The first argument will be a text file containing multiple sentences, one per line. Possibly empty
lines too. e.g.
    Hello World
    Hello CodeEval
    Output sample:
Print to stdout, each line with its words reversed, one per line. Empty lines in the input should be
ignored. Ensure that there are no trailing empty spaces on each line you print.  e.g.
    World Hello
    CodeEval Hello
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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Create a list of words
		line := scanner.Text()
		line = strings.Trim(line, " ")
		wordList := strings.Split(line, " ")
		var output string
		for _, word := range wordList {
			// Build a string in reversed-word order
			output = word + " " + output
		}
		// Remove the last character, always a space
		output = output[:len(output)-1]
		fmt.Println(output)
	}
}

/* Ported from the following Python:
import sys

def main():
    """Reverses the words of an input sentence.

    The input must be passed as the first command-line argument.

    """
    with open(sys.argv[1], 'r') as f:
        for line in f:
            # Create a list of words
            wordlist = line.split()
            output = ''
            for word in wordlist:
                # Build a string in reversed-word order
                output = str(word + ' ' + output)
            # Remove the last character, always a space
            output = output[:-1]
            print(output)

if __name__ == '__main__':
    main()
*/
