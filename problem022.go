/*
Problem 22: Fibonacci Series

Challenge Description:

The Fibonacci series is defined as:

    F(0) = 0; F(1) = 1; F(n) = F(n-1) + F(n-2) when n>1;.

Given a positive integer 'n', print out the F(n).

Input sample:

The first argument will be a text file containing a positive integer, one per line. e.g.

    5
    12

Output sample:

Print to stdout, the fibonacci number, F(n).
e.g.

    5
    144
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Avoid panic reading os.Args[1]
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./problem22 <file>")
		os.Exit(0)
	}
	// Instantiate file and scanner
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// n = int(line)
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
            // Base case
			fmt.Println(0)
		} else if n == 1 {
            // Base case
			fmt.Println(1)
		} else {
            // Fib
			a, b := 0, 1
			for i := 1; i < n; i++ {
				a, b = b, a+b
			}
			fmt.Println(b)
		}
	}
    if err := file.Close(); err != nil {
        log.Fatal(err)
    }
}

/* Portlet from the following Python:
import sys
def main():
    with open(sys.argv[1], 'r') as f:
        for line in f:
            n = int(line)
            if n == 0:
                print(0)
            if n == 1:
                print(1)
            else:
                a, b = 0, 1
                for i in range(1, n):
                    a, b = b, a + b
                print(b)

if __name__ == '__main__':
    main()
*/
