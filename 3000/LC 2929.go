func distributeCandies(n int, limit int) (ans int64) {
    for i := max(0, n - limit * 2); i <= min(n, limit); i++ {
        k := n - i
        if k <= limit {
            ans += int64(k + 1)
        } else {
            ans += int64(limit * 2 - k + 1)
        }
    }

    return
}