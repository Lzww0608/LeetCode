func countIslands(grid [][]int, k int) (ans int) {
    m, n := len(grid), len(grid[0])

    vis := make([]bool, m * n)
    var dfs func(int, int) int 
    dfs = func(i, j int) int {
        if i < 0 || i >= m || j < 0 || j >= n || vis[i * n + j] || grid[i][j] == 0 {
            return 0
        }
        vis[i * n + j] = true

        res := grid[i][j]
        res += dfs(i + 1, j)
        res += dfs(i, j + 1)
        res += dfs(i - 1, j)
        res += dfs(i, j - 1)
        return res
    }

    for i := range m {
        for j := range n {
            x := dfs(i, j)
            if x != 0 && x % k == 0 {
                ans++
            }
        }
    }

    return 
}