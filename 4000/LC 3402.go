func minimumOperations(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0])
    for j := range n {
        for i := 1; i < m; i++ {
            ans += max(0, - grid[i][j] + grid[i - 1][j] + 1)
            grid[i][j] = max(grid[i][j], grid[i - 1][j] + 1)
        }
    }

    return
}