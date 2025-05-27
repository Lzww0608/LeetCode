func wordBreak(s string, wordDict []string) bool {
    m := make(map[string]bool)
    mx := 0
    for _, t := range wordDict {
        m[t] = true
        mx = max(mx, len(t))
    }

    n := len(s)
    f := make([]bool, n + 1)
    f[0] = true
    for i := 0; i < n; i++ {
        if !f[i] {
            continue
        }

        for d := 1; d <= mx && i + d <= n; d++ {
            if m[s[i:i+d]] {
                f[i+d] = true
            }
        }
    }

    return f[n]
}