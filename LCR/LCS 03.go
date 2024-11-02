func largestArea(grid []string) (ans int) {
    n, m := len(grid), len(grid[0])
    var closed bool
    var area int
    var c byte
    vis := make([]bool, m * n)
    var dfs func(int, int)
    dfs = func(i, j int) {
        if i == -1 || i == n || j == -1 || j == m || grid[i][j] == '0' {
            closed = false
            return
        } else if grid[i][j] != c || vis[i * m + j] {
            return
        }
        area++
        vis[i * m + j] = true
        dfs(i + 1, j)
        dfs(i - 1, j)
        dfs(i, j + 1)
        dfs(i, j - 1)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] != '0' && !vis[i * m + j] {
                closed = true
                area = 0
                c = grid[i][j]
                dfs(i,j)
                if closed {
                    ans = max(ans, area)
                }
            }
        }
    }
    return 
}