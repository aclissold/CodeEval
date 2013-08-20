/*
Problem 24: Sum of Integers from File

Challenge Description:

Print out the sum of integers read from a file.

Input sample:

The first argument to the program will be a text file containing a positive integer, one per line. E.g.

    5
    12

Output sample:

Print out the sum of all the integers read from the file. E.g.

    17

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
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./problem24 <file>")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		integer, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum += integer
	}
	fmt.Println(sum)
}
