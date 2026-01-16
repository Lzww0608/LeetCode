func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
    for i := x; i < x + k / 2; i++ {
        for j := y; j < y + k; j++ {
            grid[i][j], grid[x + k - (i - x + 1)][j] = grid[x + k - (i - x + 1)][j], grid[i][j]
        } 
    }

    return grid
}