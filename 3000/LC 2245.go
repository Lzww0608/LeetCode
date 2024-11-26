func maxTrailingZeros(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0])
    type pair struct {
        x, y int
    }
    pre_row := make([][]pair, m)
    suf_row := make([][]pair, m)
    pre_col := make([][]pair, m + 1)
    suf_col := make([][]pair, m + 1)
    pre_col[m] = make([]pair, n + 1)
    suf_col[m] = make([]pair, n + 1)

    for i := 0; i < m; i++ {
        pre_row[i] = make([]pair, n + 1)
        suf_row[i] = make([]pair, n + 1)
        pre_col[i] = make([]pair, n + 1)
        suf_col[i] = make([]pair, n + 1)

        for j := 0; j < n; j++ {
            pre_row[i][j+1].x += pre_row[i][j].x + calc_two(grid[i][j])
            pre_row[i][j+1].y += pre_row[i][j].y + calc_five(grid[i][j])
        }

        for j := n - 1; j >= 0; j-- {
            suf_row[i][j].x += suf_row[i][j+1].x + calc_two(grid[i][j])
            suf_row[i][j].y += suf_row[i][j+1].y + calc_five(grid[i][j])
        }
    }

    for j := 0; j < n; j++ {
        for i := 0; i < m; i++ {
            pre_col[i+1][j].x += pre_col[i][j].x + calc_two(grid[i][j])
            pre_col[i+1][j].y += pre_col[i][j].y + calc_five(grid[i][j])
        }

        for i := m - 1; i >= 0; i-- {
            suf_col[i][j].x += suf_col[i+1][j].x + calc_two(grid[i][j])
            suf_col[i][j].y += suf_col[i+1][j].y + calc_five(grid[i][j])
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            x, y := pre_row[i][j+1].x + pre_col[i][j].x, pre_row[i][j+1].y + pre_col[i][j].y
            ans = max(ans, min(x, y))
            x, y = pre_row[i][j+1].x + suf_col[i+1][j].x, pre_row[i][j+1].y + suf_col[i+1][j].y
            ans = max(ans, min(x, y))
            x, y = suf_row[i][j].x + suf_col[i+1][j].x, suf_row[i][j].y + suf_col[i+1][j].y
            ans = max(ans, min(x, y))
            x, y = suf_row[i][j].x + pre_col[i][j].x, suf_row[i][j].y + pre_col[i][j].y
            ans = max(ans, min(x, y))
        }
    }

    return
}

func calc_five(x int) (cnt int) {
    for x % 5 == 0 {
        cnt++
        x /= 5
    }

    return 
}

func calc_two(x int) (cnt int) {
    for x % 2 == 0 {
        cnt++
        x /= 2
    }

    return 
}
