/*
Problem 26: File Size

Challenge Description:

Print the size of a file in bytes.

Input:

The first argument to your program has the path to the file you need to check the size of.

Output sample:

Print the size of the file in bytes. e.g.

    55

*/
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
    // Avoid error accessing os.Args[1]
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./problem26 <file>")
		os.Exit(1)
	}
    // Open file
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
    // Get file status
	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
    // Print filesize
	fmt.Println(fi.Size())
}
