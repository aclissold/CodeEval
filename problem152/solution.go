/*
Problem 152: Age Distribution

CHALLENGE DESCRIPTION:

You're responsible for providing a demographic report for your local school
district based on age.  To do this, you're going determine which 'category' each
person fits into based on their age.

The person's age will determine which category they should be in:

If they're from 0 to 2 the child should be with parents print : 'Still in Mama's arms'
If they're 3 or 4 and should be in preschool print : 'Preschool Maniac'
If they're from 5 to 11 and should be in elementary school print : 'Elementary school'
From 12 to 14: 'Middle school'
From 15 to 18: 'High school'
From 19 to 22: 'College'
From 23 to 65: 'Working for the man'
From 66 to 100: 'The Golden Years'
If the age of the person less than 0 or more than 100 - it might be an alien - print: "This program is for humans"

INPUT SAMPLE:

Your program should accept as its first argument a path to a filename. Each line
of input contains one integer - age of the person:

	0
	19

OUTPUT SAMPLE:

For each line of input print out where the person is:

	Still in Mama's arms
	College

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
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		age := scanner.Text()
		fmt.Println(category(age))
	}
}

// category returns a categorical description of age.
func category(age string) string {
	i, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
	}
	switch {
	case i >= 0 && i <= 2:
		return "Still in Mama's arms"
	case i >= 3 && i <= 4:
		return "Preschool Maniac"
	case i >= 5 && i <= 11:
		return "Elementary school"
	case i >= 12 && i <= 14:
		return "Middle school"
	case i >= 15 && i <= 18:
		return "High school"
	case i >= 19 && i <= 22:
		return "College"
	case i >= 23 && i <= 65:
		return "Working for the man"
	case i >= 66 && i <= 100:
		return "The Golden Years"
	default:
		return "This program is for humans"
	}
}
