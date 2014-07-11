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

for row in 1...12 {
    print(row)
    for column in 2...12 {
        let number = row * column
        var spaces: String
        switch number {
        case 100...Int.max:
            spaces = " "
        case 10..<100:
            spaces = "  "
        default:
            spaces = "   "
        }
        print("\(spaces)\(number)")
    }
    println()
}
