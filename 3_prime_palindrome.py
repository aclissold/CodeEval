# Problem 3: Prime Palindrome
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