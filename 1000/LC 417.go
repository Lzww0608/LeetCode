func pacificAtlantic(heights [][]int) (ans [][]int) {
    m, n := len(heights), len(heights[0])
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    vis := make(map[int]bool)
    vis2 := make(map[int]bool)

    var dfs func(int, int)
    dfs = func(i, j int) {
        vis[i * n + j] = true 
        t := heights[i][j]
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= m || y < 0 || y >= n || vis[x * n + y] {
                continue 
            }

            if heights[x][y] >= t {
                dfs(x, y)
            }
        }
    }
    for i := 0; i < m; i++ {
        dfs(i, 0)
    }
    for j := 0; j < n; j++ {
        dfs(0, j)
    }

    var dfs2 func(int, int)
    dfs2 = func(i, j int) {
        vis2[i * n + j] = true
        if vis[i * n + j] {
            ans = append(ans, []int{i, j})
            vis[i * n + j] = false
        }
        t := heights[i][j]
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x < 0 || x >= m || y < 0 || y >= n || vis2[x * n + y] {
                continue 
            }

            if heights[x][y] >= t {
                dfs2(x, y)
            }
        }
    }
    for i := 0; i < m; i++ {
        dfs2(i, n - 1)
    }
    for j := 0; j < n; j++ {
        dfs2(m - 1, j)
    }

    return
}