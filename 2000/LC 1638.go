func countSubstrings(s string, t string) (ans int) {
    n, m := len(s), len(t)
    // 0 <= i < n, 0 <= j < m
    // 1 - m <= i - j < n
    // d = i - j
    for d := 1 - m; d < n; d++ {
        i := max(d, 0)
        for k0, k1 := i - 1, i - 1; i < n && i - d < m; i++ {
            if s[i] != t[i-d] {
                k0 = k1
                k1 = i
            }
            ans += k1 - k0
        }
    }

    return
}