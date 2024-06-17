func maximumScore(a int, b int, c int) int {
    sum := a + b + c
    a, c = min(a, b, c), max(a, b, c)
    b = sum - a - c

    if a + b <= c {
        return a + b
    }

    return c + (a + b - c) / 2
}
