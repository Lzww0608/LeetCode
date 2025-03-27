func minimumCost(s string) int64 {
    ans := 0
    n := len(s)

    for i := 1; i < n; i++ {
        if s[i] != s[i-1] {
            ans += min(i, n - i)
        }
    }

    return int64(ans)
}