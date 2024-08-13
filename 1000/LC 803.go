func hitBricks(grid [][]int, hits [][]int) []int {
    m, n := len(grid), len(grid[0])
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    var dfs func(int, int) int
    dfs = func(i, j int) int {
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 {
            return 0
        } 
        grid[i][j] = 2
        return dfs(i + 1, j) + dfs(i - 1, j) + dfs(i, j + 1) + dfs(i, j - 1) + 1
    }

    isConnected := func(i, j int) bool {
        if i == 0 {
            return true
        }
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 2 {
                return true
            }
        }
        return false
    }

    for _, v := range hits {
        grid[v[0]][v[1]]--
    }

    for j := range grid[0] {
        dfs(0, j)
    }

    ans := make([]int, len(hits))
    for i := len(hits) - 1; i >= 0; i-- {
        v := hits[i]
        grid[v[0]][v[1]]++
        if grid[v[0]][v[1]] == 1 && isConnected(v[0], v[1]) {
            ans[i] = dfs(v[0], v[1]) - 1
        }
    }

    return ans
}