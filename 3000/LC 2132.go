func possibleToStamp(grid [][]int, h int, w int) bool {
    m, n := len(grid), len(grid[0])
    pre := make([][]int, m + 2)
    dif := make([][]int, m + 2)
    for i := 0; i <= m + 1; i++ {
        pre[i] = make([]int, n + 2)
        dif[i] = make([]int, n + 2)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            pre[i + 1][j + 1] = pre[i + 1][j] + pre[i][j + 1] - pre[i][j] + grid[i][j]
        }
    }

    for i := 0; i + h <= m; i++ {
        for j := 0; j + w <= n; j++ {
            cur := pre[i + h][j + w] - pre[i + h][j] - pre[i][j + w] + pre[i][j]
            if cur == 0 {
                dif[i + 1][j + 1] += 1
                dif[i + 1][j + w + 1] -= 1
                dif[i + h + 1][j + 1] -= 1
                dif[i + h + 1][j + w + 1] += 1 
            }
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            dif[i + 1][j + 1] += dif[i + 1][j] + dif[i][j + 1] - dif[i][j]
            if grid[i][j] == 0 && dif[i + 1][j + 1] == 0 {
                return false
            } 
        }
    }

    return true
}