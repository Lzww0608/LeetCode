func distributeCandies(n int, limit int) int64 {
    return C2(n + 2) - 3 * C2(n + 2 - (limit + 1)) + 3 * C2(n + 2 - 2 * (limit + 1)) - C2(n + 2 - 3 * (limit + 1))
}

func C2(n int) int64 {
    if n > 1 {
        return int64(n * (n - 1) / 2)
    }

    return 0
}