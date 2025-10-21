func minCost(m int, n int, waitCost [][]int) int64 {
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
        for j := range f[i] {
            f[i][j] = math.MaxInt / 2
        }
    }

    f[1][0] = -waitCost[0][0]
    for i := range m {
        for j := range n {
            f[i + 1][j + 1] = min(f[i][j + 1], f[i + 1][j]) + waitCost[i][j] + (i + 1) * (j + 1)
        }
    }

    return int64(f[m][n] - waitCost[m - 1][n - 1])
}