func numDistinct(s string, t string) int {
    n, m := len(s), len(t)
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
    }
    for j := range f[0] {
        f[0][j] = 1
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if t[i] == s[j] {
                f[i+1][j+1] = f[i][j] + f[i+1][j]
            } else {
                f[i+1][j+1] = f[i+1][j]
            }
        }
    }

    return f[m][n]
}