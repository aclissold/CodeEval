/*
Problem 25: Odd Numbers

Challenge Description:

Print the odd numbers from 1 to 99.

Input sample:

There is no input for this program.

Output sample:

Print the odd numbers from 1 to 99, one number per line. 

*/
package main

import "fmt"

func main() {
	for num := 1; num < 100; num += 2 {
		fmt.Println(num)
	}
}
