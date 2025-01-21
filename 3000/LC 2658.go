func findMaxFish(grid [][]int) (ans int) {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    m, n := len(grid), len(grid[0])
    vis := make([]bool, m * n)

    bfs := func(start_x, start_y int) (cnt int) {
        q := [][]int{{start_x, start_y}}
        for len(q) > 0 {
            cur := q[0]
            q = q[1:]
            i, j := cur[0], cur[1]
            cnt += grid[i][j]
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] > 0 && !vis[x * n + y] {
                    vis[x * n + y] = true
                    q = append(q, []int{x, y})
                }
            }
        }
        return
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] > 0 && !vis[i * n + j] {
                vis[i * n + j] = true
                ans = max(ans, bfs(i, j))
            }
        }
    }

    return
}