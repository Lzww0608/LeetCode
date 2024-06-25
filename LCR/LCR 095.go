func longestCommonSubsequence(a string, b string) int {
    m, n := len(a), len(b)
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
    }

    for i := range a {
        for j := range b {
            if a[i] == b[j] {
                f[i+1][j+1] = max(f[i+1][j+1], f[i][j] + 1)
            } else {
                f[i+1][j+1] = max(f[i+1][j], f[i][j+1])
            }
        }
    }

    return f[m][n]
}
