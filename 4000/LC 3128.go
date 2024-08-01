func numberOfRightTriangles(grid [][]int) int64 {
    n, m := len(grid), len(grid[0])
    f := make([]int, n)
    g := make([]int, m)
    for i, row := range grid {
        for j := range row {
            if grid[i][j] == 1 {
                f[i]++
                g[j]++
            }
        }
    }

    ans := 0
    for i, row := range grid {
        for j, x := range row {
            if x == 1 {
                ans += (f[i] - 1) * (g[j] - 1)
            }
        }
    }

    return int64(ans)
}