func getSum(x int, y int) int {
    for y != 0 {
        x, y = x ^ y, (x & y) << 1
    }

    return x
}