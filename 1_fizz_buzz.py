import sys

# Problem 1: Fizz Buzz
def main():
    count = 0
    # Assign the file to variable "f" as read-only
    f = open(sys.argv[1], 'r')
    for line in f:
        count += 1
        if count > 1:
            print()
        # Create a list of values for this particular line
        values = [int(value) for value in line.split()]
        for i in range(1, values[2] + 1):
            if (i % values[0] is 0) and (i % values[1] is 0):
                print('FB', end='')
            elif i % values[0] is 0:
                print('F', end='')
            elif i % values[1] is 0:
                print('B', end='')
            else:
                print(i, end='')
            if i < values[2]:
                print(' ', end='')
    f.close()

if __name__ == '__main__':
    main()

