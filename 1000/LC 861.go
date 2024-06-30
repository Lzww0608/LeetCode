func matrixScore(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    for i := range grid {
        if grid[i][0] == 0 {
            for j := 0; j < m; j++ {
                grid[i][j] ^= 1
            }
        }
    }
    ans := n * (1 << (m - 1))
    for j := 1; j < m; j++ {
        t := 0
        for i := 0; i < n; i++ {
            if grid[i][j] == 0 {
                t++
            }
        }
        for i := 0; i < n; i++ {
            if t > n / 2 {
                ans += (grid[i][j] ^ 1) << (m - j - 1)
            } else {
                ans += grid[i][j] << (m - j - 1)
            }
        }
    }
    return ans
}
