/*

Problem 29: Unique Elements

Challenge Description:

You are given a sorted list of numbers with duplicates. Print out the sorted list with duplicates
removed.

Input sample:

File containing a list of sorted integers, comma delimited, one per line. E.g. 

	1,1,1,2,2,3,3,4,4
	2,3,4,5,5

Output sample:

Print out the sorted list with duplicates removed, one per line. 
E.g.

	1,2,3,4
	2,3,4,5

*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func contains(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./problem29 <file>")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input := strings.Split(line, ",")
		var output = make([]string, 0, 10)
		for _, v := range input {
			if contains(output, v) {
				continue
			}
			output = append(output, v)
		}
		fmt.Println(strings.Join(output, ","))
	}
}
