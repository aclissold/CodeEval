/*
Problem 103: Lowest Unique Number

CHALLENGE DESCRIPTION:

There is a game where each player picks a number from 1 to 9, writes it on a
paper and gives to a guide. A player wins if his number is the lowest unique. We
may have 10-20 players in our game.

INPUT SAMPLE:

Your program should accept as its first argument a path to a filename.

You're a guide and you're given a set of numbers from players for the round of
game. E.g. 2 rounds of the game look this way:

	3 3 9 1 6 5 8 1 5 3
	9 2 9 9 1 8 8 8 2 1 1

OUTPUT SAMPLE:

Print a winner's position or 0 in case there is no winner. In the first line of
input sample the lowest unique number is 6. So player 5 wins.

	5
	0

*/
package main

import (
	"bufio"
	"fmt"
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
		fmt.Println(winner(scanner.Text()))
	}
}

const min, max int = 1, 9

type datum struct {
	position, count int
}

// winner returns the 1-indexed position of the lowest unique number in s or 0
// if there is none. s must consist of space-separated integers in [min, max].
func winner(s string) int {
	data := make(map[string]datum)
	for i := min; i <= max; i++ {
		data[strconv.Itoa(i)] = datum{position: -1, count: 0}
	}

	for i, n := range strings.Split(s, " ") {
		data[n] = datum{i + 1, data[n].count}
		data[n] = datum{data[n].position, data[n].count + 1}
	}

	for i := min; i <= max; i++ {
		a := strconv.Itoa(i)
		if data[a].count == 1 {
			return data[a].position
		}
	}

	return 0
}
