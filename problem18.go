/*
Problem 18: Multiples of a Number

Challenge Description:

Given numbers x and n, where n is a power of 2, print out the smallest
multiple of n which is greater than or equal to x. Do not use division
or modulo operator.

Input sample:

The first argument will be a text file containing a comma separated list
of two integers, one list per line. e.g.

13,8
17,16
Output sample:

Print to stdout, the smallest multiple of n which is greater than or equal
to x, one per line.
e.g.

16
32
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
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./problem18 <file>")
		os.Exit(0)
	}
	// Instantiate a file and scanner
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Retrieve x and n
		data := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(data[0])
		if err != nil {
			log.Fatal(err)
		}
		n, err := strconv.Atoi(data[1])
		if err != nil {
			log.Fatal(err)
		}
		// Initialize and increment the multiplier
		multiplier := 1
		for x > multiplier*n {
			if x > multiplier*n {
				multiplier++
			}
		}

		// Print the result
		fmt.Println(multiplier * n)
	}
}

/* Ported from the following Python:
import sys

def main():
    with open(sys.argv[1], 'r') as f:
        for line in f:

            # Retrieve x and n
            data = line.split(',')
            x, n = int(data[0]), int(data[1])

            # Initialize and increment the multiplier
            multiplier = 1
            while x > (multiplier * n):
                if x > (multiplier * n):
                    multiplier += 1

            # Print the result
            print(multiplier * n)

if __name__ == '__main__':
    main()
*/
