func numIslands(grid [][]byte) (ans int) {
    m, n := len(grid), len(grid[0])
    vis := make([]bool, m * n)

    var dfs func(int, int)
    dfs = func(i, j int) {
        if i < 0 || i >= m || j < 0 || j >= n || vis[i * n + j] || grid[i][j] != '1' {
            return
        }
        grid[i][j] = '2'
        vis[i * n + j] = true
        dfs(i + 1, j)
        dfs(i, j + 1)
        dfs(i - 1, j)
        dfs(i, j - 1)

    }

    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == '1' {
                ans++
                dfs(i, j)
            }
            
        }
    }

    return
}
