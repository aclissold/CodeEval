// Problem 4: Sum of Primes
//
// Challenge Description:
//
// Write a program to determine the sum of the first 1000 prime numbers.
//
// Input sample:
//
// None
//
// Output sample:
//
// Your program should print the sum on stdout.i.e.
//
// 3682913

package main

import "fmt"

// Prints the sum of the first 1000 prime numbers.
func main() {
	primeSum := 2 // A running summation
	primes := 1   // A count of how many primes were found
	num := 3      // The current integer to check for primeness
	isPrime := true
	for primes < 1000 {
		isPrime = true
		for i := 2; i < num; i++ {
			if num%i == 0 { // If prime
				isPrime = false
				break
			}
		}
		if isPrime {
			primeSum += num
			primes++
		}
		num++
	}
	fmt.Println(primeSum)
}
