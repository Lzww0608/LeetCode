func largestIsland(grid [][]int) (ans int) {
    n := len(grid)
    idx := make([][]int, n)
    cnt := -1
    for i := range idx {
        idx[i] = make([]int, n)
    }

    var dfs func([][]int, int, int, int) int
    dfs = func(m [][]int, i, j, cnt int) int {
        if i < 0 || i >= n || j < 0 || j >= n || m[i][j] < 0 || grid[i][j] == 0 {
            return 0
        }
        m[i][j] = cnt
        return 1 + dfs(m, i + 1, j, cnt) + dfs(m, i - 1, j, cnt) + dfs(m, i, j + 1, cnt) + dfs(m, i, j - 1, cnt)
    }

    for i := range grid {
        for j := range grid {
            if grid[i][j] == 1 {
                dfs(grid, i, j, -dfs(idx, i, j, cnt))
                cnt--
            }
        }
    }

    for i := range grid {
        for j := range grid {
            if grid[i][j] != 0 {
                ans = max(ans, -grid[i][j])
            } else {
                a, b, c, d := 0, 0, 0, 0
                cal := func(x, y int) int {
                    if x < 0 || x >= n || y < 0 || y >= n || grid[x][y] == 0 {
                        return 0
                    } 
                    t := idx[x][y]
                    if a == t || b == t || c == t || d == t {
                        return 0
                    }
                    switch {
                        case a == 0:
                            a = t
                        case b == 0:
                            b = t
                        case c == 0:
                            c = t
                        case b == 0:
                            d = t
                    } 
                    return -grid[x][y]
                }
                ans = max(ans, 1 + cal(i + 1, j) + cal(i - 1, j) + cal(i, j + 1) + cal(i, j - 1))
            }
        }
    }

    return 

}