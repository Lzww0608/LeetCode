func articleMatch(s string, p string) bool {
    m, n := len(s), len(p)
    f := make([][]bool, m + 1)
    for i := range f {
        f[i] = make([]bool, n + 1)
    }
    f[0][0] = true

    check := func(i, j int) bool {
        if i == 0 {
            return false
        } else if p[j - 1] == '.' {
            return true
        }
        return s[i-1] == p[j-1]
    }

    for i := 0; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if p[j-1] == '*' {
                f[i][j] = f[i][j-2]
                if check(i, j - 1) {
                    f[i][j] = f[i][j] || f[i-1][j]
                }
            } else if check(i, j) {
                f[i][j] = f[i-1][j-1]
            }
        }
    }

    return f[m][n]
}