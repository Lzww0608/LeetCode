func hasValidPath(grid [][]int) bool {
    m, n := len(grid), len(grid[0])
    vis := make([]bool, m * n)

    directions := map[int][][2]int{
        1: {{0, -1}, {0, 1}},    // left, right
        2: {{-1, 0}, {1, 0}},    // up, down
        3: {{0, -1}, {1, 0}},    // left, down
        4: {{0, 1}, {1, 0}},     // right, down
        5: {{0, -1}, {-1, 0}},   // left, up
        6: {{0, 1}, {-1, 0}},    // right, up
    }

    connected := func(x, y, nx, ny int) bool {
        if nx < 0 || ny < 0 || nx >= m || ny >= n {
            return false
        }
        for _, d := range directions[grid[nx][ny]] {
            if nx+d[0] == x && ny+d[1] == y {
                return true
            }
        }
        return false
    }

    var dfs func(int, int) bool
    dfs = func(x, y int) bool {
        if x == m-1 && y == n-1 {
            return true
        }
        vis[x*n+y] = true
        for _, d := range directions[grid[x][y]] {
            nx, ny := x+d[0], y+d[1]
            if nx >= 0 && ny >= 0 && nx < m && ny < n && !vis[nx*n+ny] && connected(x, y, nx, ny) {
                if dfs(nx, ny) {
                    return true
                }
            }
        }
        return false
    }

    return dfs(0, 0)
}
