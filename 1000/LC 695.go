var dir = []int{1, 0, -1, 0}
func maxAreaOfIsland(grid [][]int) int {
    ans, cnt := 0, 0
    n, m := len(grid), len(grid[0])
    var dfs func(i, j int)
    dfs = func(i, j int) {
        if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] == 0 {
            return 
        }
        grid[i][j] = 0
        cnt++
        ans = max(ans, cnt)
        for k := 0; k < 4; k++ {
            x, y := i + dir[k], j + dir[(k + 1) % 4]
            dfs(x, y)
        }
    }
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 1 {
                cnt = 0
                dfs(i, j)
            }
        }
    }
    return ans
}
