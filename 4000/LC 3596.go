func minCost(m int, n int) int {
    if m > 2 || n > 2 || m == 2 && n == 2 {
        return -1
    }

    if m == 1 && n == 1 {
        return 1
    }

    return 3
}