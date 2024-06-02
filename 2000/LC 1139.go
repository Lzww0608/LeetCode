func largest1BorderedSquare(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    //ans := 0
    row, col := make([][]int, n+1), make([][]int, n+1)
    for i := range row {
        row[i], col[i] = make([]int, m+1), make([]int, m+1)
    }
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 1 {
                row[i][j+1] = row[i][j] + 1
                col[i+1][j] = col[i][j] + 1
            } else {
                row[i][j+1] = row[i][j]
                col[i+1][j] = col[i][j]
            }
        }
    }

    check := func(i, j, k int) bool {
        if row[i][j+k] - row[i][j] == k && row[i+k-1][j+k] - row[i+k-1][j] == k && col[i+k][j] - col[i][j] == k && col[i+k][j+k-1] - col[i][j+k-1] == k {
            return true
        }
        return false
    }

    for k := min(n, m); k > 0; k-- {
        for i := 0; i < n && (n - i) >= k; i++ {
            for j := 0; j < m && (m - j) >= k; j++ {
                if check(i, j, k) {
                    return k * k
                }
            }
        }
    }
    return 0
}
