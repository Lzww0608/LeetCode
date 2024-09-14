func tilingRectangle(n int, m int) int {
    g := make([][]int, n)
    for i := range g {
        g[i] = make([]int, m)
    }

    fill := func(x1, y1, x2, y2, val int) {
        for i := x1; i <= x2; i++ {
            for j := y1; j <= y2; j++ {
                g[i][j] += val
            }
        }
    }

    ans := n * m 
    var dfs func(int, int, int)
    dfs = func(i, j, cnt int) {
        if cnt >= ans {
            return
        }

        if i == n {
            ans = min(ans, cnt)
            return
        }

        if g[i][j] == 1 {
            if j + 1 < m {
                dfs(i, j + 1, cnt)
            } else {
                dfs(i + 1, 0, cnt)
            }
        } 

        d := 1
        for i + d - 1 < n && j + d - 1 < m && g[i][j+d-1] == 0 {
            d++
        }
        d--
        for k := d; k >= 1; k-- {
            fill(i, j, i + k - 1, j + k - 1, 1)
            if j + k < m {
                dfs(i, j + k, cnt + 1)
            } else {
                dfs(i + 1, 0, cnt + 1)
            }
            
            fill(i, j, i + k - 1, j + k - 1, -1)
        }
    }

    dfs(0, 0, 0)

    return ans
}