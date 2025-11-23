func numberOfChild(n int, k int) int {
    t := k % (n - 1)
    if (k / (n - 1)) & 1 == 0 {
        return t
    }

    return n - 1 - t
}