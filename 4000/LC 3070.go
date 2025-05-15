func countSubmatrices(grid [][]int, k int) (ans int) {
    m, n := len(grid), len(grid[0])
    pre := make([][]int, m + 1)
    for i := range pre {
        pre[i] = make([]int, n + 1)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            pre[i+1][j+1] = pre[i][j+1] + pre[i+1][j] - pre[i][j] + grid[i][j] 
            if pre[i+1][j+1] <= k {
                ans++
            }
        }
    }

    return
}