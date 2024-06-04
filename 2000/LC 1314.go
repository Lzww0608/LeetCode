func matrixBlockSum(mat [][]int, k int) [][]int {
    n, m := len(mat), len(mat[0])
    ans, sum := make([][]int, n), make([][]int, n+1)
    for i := range sum {
        sum[i] = make([]int, m+1)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            sum[i+1][j+1] = mat[i][j] + sum[i][j+1] + sum[i+1][j] - sum[i][j]
        }
    }
    for i := range ans {
        ans[i] = make([]int, m)
        for j := 0; j < m; j++ {
            x1, y1, x2, y2 := max(0, i-k), max(0,j-k), min(n,i+k+1), min(m,j+k+1)
            ans[i][j] = sum[x2][y2] - sum[x2][y1] - sum[x1][y2] + sum[x1][y1]
        }
    }
    return ans
}
