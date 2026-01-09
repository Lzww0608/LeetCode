func minCuttingCost(n int, m int, k int) int64 {
    if n <= k && m <= k {
        return 0
    } 

    if n > k {
        return int64(k * (n - k))
    }

    return int64(k * (m - k))
}