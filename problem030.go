/*
Problem 30: Set Intersection

Challenge Description:

You are given two sorted list of numbers (ascending order). The lists themselves are comma delimited
and the two lists are semicolon delimited. Print out the intersection of these two sets.

Input sample:

File containing two lists of ascending order sorted integers, comma delimited, one per line. E.g.

	1,2,3,4;4,5,6
	20,21,22;45,46,47
	7,8,9;8,9,10,11,12

Output sample:

Print out the ascending order sorted intersection of the two lists, one per line. Print empty new
line in case the lists have no intersection. E.g.

	4

	8,9

*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// contains returns whether or not a slice "a" contains a string "s".
func contains(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

func main() {
	// Read the file passed as a command line argument into data
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// Parse individual lines from the file into a slice
	lines := strings.Split(string(data), "\n")
	// Remove the empty string that follows the final \n
	lines = lines[:len(lines)-1]

	for _, line := range lines {
		// Split the line into two lists
		lists := strings.Split(line, ";")
		list1 := strings.Split(lists[0], ",")
		list2 := strings.Split(lists[1], ",")

		intersection := make([]string, 0, 5)
		// Append all elements found in both list1 and list2 to intersection
		for _, v := range list1 {
			if contains(list2, v) {
				intersection = append(intersection, v)
			}
		}

		// Print the results
		fmt.Println(strings.Join(intersection, ","))
	}
}
