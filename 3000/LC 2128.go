func removeOnes(grid [][]int) bool {
    m, n := len(grid), len(grid[0])
    for i := 1; i < m; i++ {
        f := grid[i][0] == grid[0][0]
        for j := 1; j < n; j++ {
            if f != (grid[i][j] == grid[0][j]) {
                return false
            }
        }
    }

    return true
}