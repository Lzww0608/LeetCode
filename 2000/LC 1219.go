func getMaximumGold(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0])
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    var dfs func(i, j int) int 
    dfs = func(i, j int) int {
        tmp := grid[i][j]
        grid[i][j] = 0
        sum := 0
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
                continue
            }
            sum = max(sum, dfs(x, y))
        }
        ans = max(ans, sum + tmp)
        grid[i][j] = tmp
        return sum + tmp
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] != 0 {
                dfs(i, j)
            }
        }
    }

    return
}