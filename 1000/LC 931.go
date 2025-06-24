func minFallingPathSum(matrix [][]int) int {
    n := len(matrix)
    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int ,n)
    }
    f[0] = matrix[0]

    for i := 1; i < n; i++ {
        for j := 0; j < n; j++ {
            f[i][j] = min(f[i-1][j], f[i-1][max(j - 1, 0)], f[i-1][min(j + 1, n - 1)]) + matrix[i][j]
        }
    }

    return slices.Min(f[n-1])
}