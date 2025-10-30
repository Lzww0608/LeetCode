func numberOfCuts(n int) int {
    if n & 1 == 0 || n == 1 {
        return n / 2
    }

    return n
}