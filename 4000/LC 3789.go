func minimumCost(a int, b int, c int, x int, y int) int64 {
    if a + b <= c {
        return int64(a * x + b * y)
    }

    if x >= y {
        return int64(y * c + (x - y) * min(a, c))
    }
    return int64(x * c + (y - x) * min(b, c))
}