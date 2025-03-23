func countPyramids(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0])

    solve := func() {
        f := make([][]int, m + 1)
        for i := range f {
            f[i] = make([]int, n)
        }

        for i := m - 1; i >= 0; i-- {
            f[i][0], f[i][n-1] = grid[i][0], grid[i][n-1]
            for j := 1; j < n - 1; j++ {
                if grid[i][j] == 0 {
                    f[i][j] = 0
                } else {
                    f[i][j] = min(f[i+1][j-1], f[i+1][j], f[i+1][j+1]) + 1
                    ans += f[i][j] - 1
                }
            }
        }
    }

    solve()
    for i := 0; i < m / 2; i++ {
        grid[i], grid[m-i-1] = grid[m-i-1], grid[i]
    }
    solve()

    return 
}