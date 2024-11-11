func minDays(grid [][]int) int {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
    m, n := len(grid), len(grid[0])
    vis := map[int]bool{}
    sum := 0

    var dfs func(int, int) bool
    dfs = func(i, j int) bool {
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 || vis[i * n + j] {
            return false
        }
        vis[i * n + j] = true
        sum++
        dfs(i + 1, j)
        dfs(i - 1, j)
        dfs(i, j + 1)
        dfs(i, j - 1)
        return true
    }

    cnt := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++{
            if grid[i][j] == 1 && !vis[i * n + j] {
                if cnt > 0 {
                    return 0
                }
                cnt++
                dfs(i, j)
            }
        }
    }

    if sum < 3 {
        return sum
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                clear(vis)
                grid[i][j] = 0
                cnt := 0
                for _, dir := range dirs {
                    x, y := i + dir[0], j + dir[1]
                    if dfs(x, y) {
                        cnt++
                    }
                }
                if cnt > 1 {
                    return 1
                }
                grid[i][j] = 1
            }
        }
    }

    return 2
}