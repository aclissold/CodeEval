/*
Problem 20: Lowercase

Challenge Description:

Given a string write a program to convert it into lowercase.

Input sample:

The first argument will be a text file containing sentences, one per line. You can assume all characters are from the english language. e.g. 

HELLO CODEEVAL
This is some text
Output sample:

Print to stdout, the lowercase version of the sentence, each on a new line.
e.g.

hello codeeval
this is some text
*/
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(strings.ToLower(scanner.Text()))
    }
}

/* Ported from the following Python:
import sys

def main():
    with open(sys.argv[1], 'r') as f:
        for line in f:
            print(line.lower(), end='')

if __name__ == '__main__':
    main()
*/
