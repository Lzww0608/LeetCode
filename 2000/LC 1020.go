func numEnclaves(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0])
    vis := make([]bool, m * n)

    var dfs func(int, int)
    dfs = func(i, j int) {
        if i < 0 || i >= m || j < 0 || j >= n ||
             vis[i * n + j] || grid[i][j] != 1 {
            return
        }
        vis[i * n + j] = true 
        dfs(i + 1, j)
        dfs(i - 1, j)
        dfs(i, j + 1)
        dfs(i, j - 1)
    }
    for i := 0; i < m; i++ {
        dfs(i, 0)
        dfs(i, n - 1)
    }
    for j := 0; j < n; j++ {
        dfs(0, j)
        dfs(m - 1, j)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 && !vis[i * n + j] {
                ans++
            }
        }
    }

    return
}