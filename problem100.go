package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(even(scanner.Text()))
	}
}

// even returns "1" if the integer representation of s is even and "0" if odd.
func even(s string) string {
	i, _ := strconv.Atoi(s)
	even := i % 2 == 0
	if even {
		return "1"
	}
	return "0"
}
