func minPathSum(grid [][]int) int {
    n := len(grid[0])
    f := make([]int, n + 1)
    for i := 0; i < n; i++ {
        f[i+1] = grid[0][i] + f[i]
    }
    for i := 1; i < len(grid); i++ {
        f[0] = f[1]
        for j, x := range grid[i] {
            f[j+1] = min(f[j+1], f[j]) + x
        }
    }

    return f[n]
}
