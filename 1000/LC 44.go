func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    f := make([][]bool, m + 1)
    for i := range f {
        f[i] = make([]bool, n + 1)
    }
    f[0][0] = true
    for i := 0; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if p[j-1] == '*' {
                f[i][j] = f[i][j-1] || (i > 0 && f[i-1][j])
            } else {
                f[i][j] = i > 0 && f[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '?')
            }
        }
    }

    return f[m][n]
}

