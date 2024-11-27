func sumScores(s string) int64 {
    n := len(s)
    ans := 0
    z := make([]int, n)
    for i, l, r := 0, 0, 0; i < n; i++ {
        z[i] = max(min(z[i-l], r - i + 1), 0)
        for i + z[i] < n && s[i+z[i]] == s[z[i]] {
            l, r = i, i + z[i]
            z[i]++
        }
        ans += z[i]
    }

    return int64(ans)
}