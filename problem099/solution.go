/*
Problem 99: Calculate Distance

CHALLENGE DESCRIPTION:

You have coordinates of 2 points and need to find the distance between them.

INPUT SAMPLE:

Your program should accept as its first argument a path to a filename. Input example is the following

	(25, 4) (1, -6)
	(47, 43) (-25, -11)

All numbers in input are integers between -100 and 100.

OUTPUT SAMPLE:

Print results in the following way.

	26
	90

You don't need to round the results you receive.
They must be integer numbers between -100 and 100.
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		p1, p2 := parsePoints(scanner.Text())
		fmt.Println(d(p1, p2))
	}
}

type point struct {
	x, y int
}

// parsePoints converts a string representation of two points.
// For example: "(25, 4) (1, -6)" -> point{25, 4}, point{1, -6}
func parsePoints(s string) (p1, p2 point) {
	i := strings.Index(s, ") (") + 1
	p1 = atop(s[:i])
	p2 = atop(s[i+1:])
	return
}

// atop converts a, formatted as "(x, y)", to a point, analogous to strconv.Atoi().
func atop(a string) point {
	items := strings.Split(a, ", ")
	x, _ := strconv.Atoi(items[0][1:])
	y, _ := strconv.Atoi(items[1][:len(items[1])-1])
	return point{x, y}
}

// d returns the distance between p1 and p2.
func d(p1, p2 point) int {
	a, b := float64(p1.x-p2.x), float64(p1.y-p2.y)
	// Pythagorean theorem
	return int(math.Abs(math.Sqrt(math.Pow(a, 2.0) + math.Pow(b, 2.0))))
}
