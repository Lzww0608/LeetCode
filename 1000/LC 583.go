func minDistance(s string, t string) int {
    m, n := len(s), len(t)
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if s[i] == t[j] {
                f[i+1][j+1] = f[i][j] + 1
            } else {
                f[i+1][j+1] = max(f[i+1][j], f[i][j+1])
            }
        }
    }

    return m + n - f[m][n] * 2
}