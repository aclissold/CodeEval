// Multiplication Tables
//
// Challenge Description:
//
// Print out the grade school multiplication table upto 12*12.
//
// Input sample:
//
// None
//
// Output sample:
//
// Print out the table in a matrix like fashion, each number formatted to a
// width of 4 (The numbers are right-aligned and strip out leadeing/trailing
// spaces on each line). The first 3 line will look like:
//
// 1   2   3   4   5   6   7   8   9  10  11  12
// 2   4   6   8  10  12  14  16  18  20  22  24
// 3   6   9  12  15  18  21  24  27  30  33  36

for num in 1..13 {
    print(num)
    for multiplier in 2..13 {
        let out = num * multiplier
        var spaces = "   "
        if out > 9 {
            if out < 100 {
                spaces = "  "
            } else {
                spaces = " "
            }
        }
        print("\(spaces)\(out)")
    }
    println()
}
