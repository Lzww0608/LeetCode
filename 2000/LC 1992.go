func findFarmland(land [][]int) (ans [][]int) {
    m, n := len(land), len(land[0])
    vis := make([]bool, m * n)
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    var start_x, start_y, end_x, end_y int
    var dfs func(int, int) int
    dfs = func(i, j int) (cnt int) {
        if i < 0 || i >= m || j < 0 || j >= n || land[i][j] != 1 || vis[i * n + j] {
            return
        }
        vis[i * n + j] = true
        start_x = min(start_x, i)
        end_x = max(end_x, i)
        start_y = min(start_y, j)
        end_y = max(end_y, j)
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            cnt += dfs(x, y)
        }

        return cnt + 1
    }

    for i, row := range land {
        for j, x := range row {
            if x == 1 && !vis[i * n + j] {
                start_x, start_y, end_x, end_y = i, j, i, j
                if dfs(i, j) == (end_x - start_x + 1) * (end_y - start_y + 1) {
                    ans = append(ans, []int{start_x, start_y, end_x, end_y})
                } 
            }
        }
    }

    return
}