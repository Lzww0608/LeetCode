func numDistinctIslands(grid [][]int) (ans int) {
    m, n := len(grid), len(grid[0])
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    island := map[string]bool{}
    seq := []byte{}

    var dfs func(int, int)
    dfs = func(i, j int) {
        for k, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
                grid[x][y] = 0
                seq = append(seq, byte(k + '0'))
                dfs(x, y)
            } else {
                seq = append(seq, byte('-'))
            }
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                grid[i][j] = 0
                seq = seq[:0]
                dfs(i, j)
                if !island[string(seq)] {
                    ans++
                    island[string(seq)] = true
                } 
            }
        }
    }

    return 
}   