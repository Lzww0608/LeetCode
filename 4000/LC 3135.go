func minOperations(s string, t string) int {
    m, n := len(s), len(t)
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
    }

    ans := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if s[i] == t[j] {
                if i > 0 && j > 0 && s[i - 1] == t[j - 1] {
                    f[i + 1][j + 1] = f[i][j] + 1 
                } else {
                    f[i + 1][j + 1] = 1
                }
            } 
            ans = max(ans, f[i + 1][j + 1])
        }
    }

    return m - ans + n - ans
}