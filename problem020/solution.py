#!/usr/bin/python3
"""Problem 20: Lowercase

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

"""
import sys

def main():
    with open(sys.argv[1], 'r') as f:
        for line in f:
            print(line.lower(), end='')

if __name__ == '__main__':
    main()
