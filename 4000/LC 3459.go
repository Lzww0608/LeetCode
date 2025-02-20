func lenOfVDiagonal(grid [][]int) (ans int) {
    dirs := [][]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
    m, n := len(grid), len(grid[0])
    memo := make([][4][2]int, m  * n)

    var dfs func(int, int, int, int, int) int
    dfs = func(i, j, k, f, target int) (res int) {
        i += dirs[k][0]
        j += dirs[k][1]
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != target {
            return 
        }
        if memo[i * n + j][k][f] != 0 {
            return memo[i * n + j][k][f]
        }
        res = dfs(i, j, k, f, 2 - target) + 1
        if f == 1 {
            res = max(res, dfs(i, j, (k + 1) % 4, 0, 2 - target) + 1)
        }
        memo[i * n + j][k][f] = res
        return
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                for k := range 4 {
                    ans = max(ans, 1 + dfs(i, j, k, 1, 2))
                }
            }
        }
    }

    return
}