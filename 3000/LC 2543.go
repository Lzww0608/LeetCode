func isReachable(x int, y int) bool {
    for x & 1 == 0 {
        x >>= 1
    }
    for y & 1 == 0 {
        y >>= 1
    }

    if x == 1 && y == 1 {
        return true
    }
    if x == y {
        return false
    }

    return isReachable(min(x, y), x + y)
}