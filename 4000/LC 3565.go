func findPath(grid [][]int, k int) (ans [][]int) {
    m, n := len(grid), len(grid[0])
    vis := make(map[int]bool)
    dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

    var dfs func(i, j, cur int) bool
    dfs = func(i, j, cur int) bool {
        ans = append(ans, []int{i, j})
        vis[i * n + j] = true 
        if grid[i][j] != 0 {
            if grid[i][j] != cur {
                ans = ans[:len(ans) - 1]
                vis[i * n + j] = false
                return false
            }
            cur++
        }

        if len(ans) == m * n {
            return cur == k + 1
        }

        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= m || y < 0 || y >= n || vis[x * n + y] {
                continue
            }
            if dfs(x, y, cur) {
                return true
            }
        }

        ans = ans[:len(ans) - 1]
        vis[i * n + j] = false
        return false
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            clear(vis)
            if dfs(i, j, 1) {
                return 
            }
        }
    }

    return 
}