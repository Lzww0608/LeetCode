func cherryPickup(grid [][]int) int {
    n := len(grid)
    f := make([][][]int, n + 1)
    for i := range f {
        f[i] = make([][]int, n + 1)
        for j := range f[i] {
            f[i][j] = make([]int, 2 * n + 1)
            for k := range f[i][j] {
                f[i][j][k] = math.MinInt32
            }
        }
    }

    for k := 2; k <= 2 * n; k++ {
        for x1 := 1; x1 <= n; x1++ {
            for x2 := 1; x2 <= n; x2++ {
                y1, y2 := k - x1, k - x2
                if y1 >= 1 && y1 <= n && y2 <= n && y2 >= 1 && grid[x1-1][y1-1] != -1 && grid[x2-1][y2-1] != -1 {
                    v := grid[x1-1][y1-1]
                    if x1 != x2 {
                        v += grid[x2-1][y2-1]
                    }
                    if k == 2 {
                        f[x1][x2][k] = v
                    } else {
                        f[x1][x2][k] = max(f[x1][x2][k], f[x1-1][x2][k-1] + v, f[x1][x2-1][k-1] + v, f[x1-1][x2-1][k-1] + v, f[x1][x2][k-1] + v)
                    }
                }
            }
        }
    }

    return max(0, f[n][n][2*n])
}
