func maximumValue(n int, s int, m int) int64 {
    ans := s
    n -= 1

    return int64(ans + m * ((n + 1) / 2) - (n - 1) / 2)
}