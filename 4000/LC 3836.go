func maxScore(a []int, b []int, k int) int64 {
    n, m := len(a), len(b)
    f := make([][][]int, k + 1)
    for i := range f {
        f[i] = make([][]int, n + 1)
        for j := range f[i] {
            f[i][j] = make([]int, m + 1)
            for t := range f[i][j] {
                f[i][j][t] = math.MinInt / 2
            }
        }
    }

    for i := range n + 1 {
        for j := range m + 1 {
            f[0][i][j] = 0
        }
    }

    for t := 1; t <= k; t++ {
        for i := t - 1; i < n; i++ {
            for j := t - 1; j < m; j++ {
                f[t][i + 1][j + 1] = max(f[t][i + 1][j], f[t][i][j + 1], f[t - 1][i][j] + a[i] * b[j])
            }
        }
    }

    return int64(f[k][n][m])
}