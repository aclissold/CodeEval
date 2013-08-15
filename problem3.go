// Problem 3: Prime Palindrome
package main

import (
	"fmt"
	"strconv"
)

func reverse(s string) string {
	// Get Unicode code points
	n := 0
	runes := make([]rune, len(s))
	for _, r := range s {
		runes[n] = r
		n++
	}
	runes = runes[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	// Convert back to a string (z = s reversed, ha)
	z := string(runes)
	return z
}

func main() {
	var isPrime bool
	var largestPalindromicPrime uint16
	var num uint16
	var i uint16
	for num = 1; num < 1000; num++ {
		isPrime = true
		for i = 2; i < num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			if strconv.Itoa(int(num)) == reverse(strconv.Itoa(int(num))) {
				largestPalindromicPrime = num
			}
		}
	}
	fmt.Print(largestPalindromicPrime)
}

/* Ported from the following Python:
def main():
    for num in range(1, 1000):
        for i in range(2, num):
            if num % i == 0:
                break
        else:
            if str(num) == str(num)[::-1]:
                largestPalindromicPrime = num
    print(largestPalindromicPrime, end='')

if __name__ == '__main__':
    main()
*/
