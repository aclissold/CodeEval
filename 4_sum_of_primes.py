#!usr/bin/python3
"""Problem 4: Sum of Primes

Challenge Description:

Write a program to determine the sum of the first 1000 prime numbers.

Input sample:

None

Output sample:

Your program should print the sum on stdout.i.e.

3682913

"""
def main():
    """Prints the sum of the first 1000 prime numbers."""
    prime_sum = 2 # A running summation
    primes = 1 # A count of how many primes were found
    num = 3 # The current integer to check for primeness
    while primes < 1000:
        for i in range(2, num):
            if num % i == 0: # If prime
                break
        else: # If no divisors found
            prime_sum += num
            primes += 1
        num += 1
    print(prime_sum)
        

if __name__ == '__main__':
    main()
