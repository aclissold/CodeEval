/*
Problem 21: Sum of Digits

Challenge Description:

Given a positive integer, find the sum of its constituent digits.

Input sample:

The first argument will be a text file containing positive integers, one per line. E.g.

    23
    496

Output sample:

Print to stdout, the sum of the numbers that make up the integer, one per line. E.g.

    5
    19

*/
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: ./problem21 <file>")
        os.Exit(0)
    }
    file, err := os.Open(os.Args[1])
    check(err)
    scanner := bufio.NewScanner(file)
    // For line in file:
    for scanner.Scan() {
        line := scanner.Text()
        var sum int
        // For character in line...
        for i := range line {
            // ...append the integer representation of that character
            digit, err := strconv.Atoi((string(line[i])))
            check(err)
            sum += digit
        }
        fmt.Println(sum)
    }
}
