func escapeMaze(g [][]string) bool {
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    k, m, n := len(g), len(g[0]), len(g[0][0])
    vis := make([][][][6]bool, k)
    for i := range vis {
        vis[i] = make([][][6]bool, m)
        for j := range vis[i] {
            vis[i][j] = make([][6]bool, n)
        }
    }

    var dfs func(int, int, int, int) bool
    dfs = func(t, x, y, s int) bool {
        if x < 0 || x >= m || y < 0 || y >= n || t + m - x - 1 + n - y - 1 >= k || vis[t][x][y][s] {
            return false
        }

        if x == m - 1 && y == n - 1 {
            return true
        }
        vis[t][x][y][s] = true

        if (s >> 1) == 1 {
            for _, dir := range dirs {
                if dfs(t + 1, x + dir[0], y + dir[1], s ^ 6) {
                    return true
                }
            }
            return dfs(t + 1, x, y, s)
        }

        if (s >> 1) == 0 && g[t][x][y] == '#' && dfs(t, x, y, s | 2) {
            return true
        }

        if g[t][x][y] == '#' {
            if s & 1 == 1 {
                return false
            }
            s |= 1
        }

        for _, dir := range dirs {
            if dfs(t + 1, x + dir[0], y + dir[1], s) {
                return true
            }
        }

        return dfs(t + 1, x, y, s)
    }

    return dfs(0, 0, 0, 0)
}