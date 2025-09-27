func leastMinutes(n int) int {
    return 1 + bits.Len(uint(n - 1))
}