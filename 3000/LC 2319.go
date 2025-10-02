func checkXMatrix(grid [][]int) bool {
    n := len(grid)
    for i := range n {
        for j := range n {
            if (i == j || i == n - j - 1) != (grid[i][j] != 0) {
                return false
            }
        }
    }

    return true
}