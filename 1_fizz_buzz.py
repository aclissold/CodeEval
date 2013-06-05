import sys

# Problem 1: Fizz Buzz
def main():
    # Assign the file to variable "f" as read-only
    f = open(sys.argv[1], 'r')
    for line in f:
        # Create a list of values for this particular line
        values = [int(x) for x in line.split()]
        for value in values:
            if (value % a is 0) and (value % b is 0):
                print('FB', end='')
            elif value % a is 0:
                print('F', end='')
            elif i % b is 0:
                print('B', end='')
            else:
                print(i, end='')
            if i < n:
                print(' ', end='')
        print()


if __name__ == '__main__':
    main()

