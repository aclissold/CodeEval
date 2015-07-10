/*
Problem 70: Overlapping Rectangles

CHALLENGE DESCRIPTION:

Given two axis aligned rectangles A and B, determine if the two overlap. The
rectangles considered overlapping if they have at least one common point.

INPUT SAMPLE:

Your program should accept as its first argument a path to a filename. Each line
in this file contains 8 comma separated co-ordinates. The co-ordinates are upper
left x of A, upper left y of A, lower right x of A, lower right y of A, upper
left x of B, upper left y of B, lower right x of B, lower right y of B. E.g.

	-3,3,-1,1,1,-1,3,-3
	-3,3,-1,1,-2,4,2,2

OUTPUT SAMPLE:

Print out True or False if A and B intersect. E.g.

	False
	True

*/
package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(overlaps(scanner.Text()))
	}
}

type rectangle struct {
	image.Rectangle
}

// touches is like image.Rectangle.Overlaps but also returns true if r and s
// have at least one partially overlapping edge.
func (r rectangle) touches(s rectangle) bool {
	return r.Min.X <= s.Max.X && s.Min.X <= r.Max.X &&
		r.Min.Y <= s.Max.Y && s.Min.Y <= r.Max.Y
}

func overlaps(coordinates string) string {
	c := strings.Split(coordinates, ",")
	ax0, _ := strconv.Atoi(c[0])
	ay0, _ := strconv.Atoi(c[1])
	ax1, _ := strconv.Atoi(c[2])
	ay1, _ := strconv.Atoi(c[3])
	bx0, _ := strconv.Atoi(c[4])
	by0, _ := strconv.Atoi(c[5])
	bx1, _ := strconv.Atoi(c[6])
	by1, _ := strconv.Atoi(c[7])

	a := rectangle{image.Rect(ax0, ay0, ax1, ay1)}
	b := rectangle{image.Rect(bx0, by0, bx1, by1)}
	if a.touches(b) {
		return "True"
	}
	return "False"
}
