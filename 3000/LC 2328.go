const MOD int = 1_000_000_007
func countPaths(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0])
    f := make([][]int, m)
    for i := range f {
        f[i] = make([]int, n)
        for j := range f[i] {
            f[i][j] = -1
        }
    }

    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    var dfs func(int, int) int 
    dfs = func(i, j int) int {
        if f[i][j] != -1 {
            return f[i][j]
        }
        res := 1
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1] 
            if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] > grid[i][j] {
                res += dfs(x, y)
            }
        }

        f[i][j] = res % MOD
        return res % MOD
    }

    for i, row := range grid {
        for j := range row {
            ans = (ans + dfs(i, j)) % MOD
        }
    }

    return
}