func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    f := make([]int, n)
    f[0] = grid[0][0]
    for i := 1; i < n; i++{
        f[i] = f[i-1] + grid[0][i]
    }

    for i := 1; i < m; i++ {
        for j := range grid[i] {
            if j > 0 {
                f[j] = min(f[j], f[j-1]) + grid[i][j]
            } else {
                f[j] += grid[i][j]  
            }
        }
    }

    return f[n-1]
}
