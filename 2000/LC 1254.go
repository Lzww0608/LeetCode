func closedIsland(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0])
    vis := make([]bool, m * n)

    var dfs func(int, int) bool
    dfs = func(i, j int) bool {
        if i < 0 || i >= m || j < 0 || j >= n {
            return false 
        } else if vis[i * n + j] || grid[i][j] == 1 {
            return true 
        }

        vis[i * n + j] = true
        ans := dfs(i + 1, j)
        ans = dfs(i - 1, j) && ans
        ans = dfs(i, j + 1) && ans
        ans = dfs(i, j - 1) && ans
        return ans
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 && !vis[i * n + j] {
                if dfs(i, j) {
                    ans++
                }
            }
        }
    }

    return 
}