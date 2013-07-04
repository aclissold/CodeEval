"""Problem 22: Fibonacci Series

Challenge Description:

The Fibonacci series is defined as:

    F(0) = 0; F(1) = 1; F(n) = F(n-1) + F(n-2) when n>1;.
    
Given a positive integer 'n', print out the F(n).



Input sample:

The first argument will be a text file containing a positive integer, one per line. e.g. 

5
12



Output sample:

Print to stdout, the fibonacci number, F(n).
e.g.

    5
    144

"""
import sys
def main():
    with open(sys.argv[1], 'r') as f:
        for line in f:
            n = int(line)
            if n == 0:
                print(0)
            if n == 1:
                print(1)
            else:
                a, b = 0, 1
                for i in range(1, n):
                    a, b = b, a + b
                print(b)
 
if __name__ == '__main__':
    main()
