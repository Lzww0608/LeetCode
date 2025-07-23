func shortestBridge(grid [][]int) int {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    n := len(grid[0])
    
    var dfs func(i, j int)
    dfs = func(i, j int) {
        if i < 0 || i >= n || j < 0 || 
            j >= n || grid[i][j] != 1 {
                return
        }

        grid[i][j] = 2 
        dfs(i + 1, j)
        dfs(i - 1, j)
        dfs(i, j - 1)
        dfs(i, j + 1)
    }

    q := []int{}
    f := false
    vis := make([]bool, n * n)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                if !f {
                    dfs(i, j)
                    f = true
                } else {
                    q = append(q, i * n + j)
                    vis[i * n + j] = true
                }
            }
        }
    }

    d := -1
    for len(q) > 0 {
        for sz := len(q); sz > 0; sz-- {
            cur := q[0]
            q = q[1:]
            i, j := cur / n, cur % n 
            if grid[i][j] == 2 {
                return d
            }
            for _, dir := range dirs {
                x, y := i + dir[0], j + dir[1]
                if x < 0 || x >= n || y < 0 || y >= n ||
                    vis[x * n + y] || grid[x][y] == 1 {
                        continue
                    }
                vis[x * n + y] = true
                q = append(q, x * n + y)
            }
        }
        d++
    }

    return -1
}