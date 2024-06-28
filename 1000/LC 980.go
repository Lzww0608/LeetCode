func uniquePathsIII(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    
    check := func() bool {
        for i := range grid {
            for _, x := range grid[i] {
                if x == 0 {
                    return false
                }
            }
        }

        return true
    }

    var dfs func(int, int) int
    dfs = func(i, j int) int {
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == -1 {
            return 0
        }
        if grid[i][j] == 2 {
            if check() {
                return 1
            }
            return 0
        }
        grid[i][j] = -1
        res := dfs(i + 1, j) + dfs(i, j + 1) + dfs(i - 1, j) + dfs(i, j - 1)
        grid[i][j] = 0
        return res
    }

    for i := range grid {
        for j, x := range grid[i] {
            if x == 1 {
                return dfs(i, j)
            }
        }
    }

    return 0
}
