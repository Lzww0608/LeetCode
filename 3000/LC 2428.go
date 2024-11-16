func maxSum(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0])

    for i := 0; i < m - 2; i++ {
        for j := 0; j < n - 2; j++ {
            ans = max(ans, grid[i][j] + grid[i][j+1] + grid[i][j+2] + grid[i+1][j+1] + grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2])
        }
    }

    return
}