func ballGame(num int, g []string) (ans [][]int) {
    m, n := len(g), len(g[0])
    dirs := [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}

    bfs := func(i, j, k int) bool {
        vis := map[[2]int]bool{}
        for p := 0; p < num; p++ {
            vis[[2]int{i * n + j, k}] = true
            i, j = i + dirs[k][0], j + dirs[k][1]
            if i < 0 || i >= m || j < 0 || j >= n {
                return false
            }
            if g[i][j] == 'O' {
                return true
            } else if g[i][j] == 'W' {
                k = (k + 3) % 4
            } else if g[i][j] == 'E' {
                k = (k + 1) % 4
            }

            if vis[[2]int{i * n + j, k}] {
                return false
            }

        }

        return false
    }
        
    

    
    for j := 1; j < n - 1; j++ {
        if g[0][j] == '.' && bfs(0, j, 0) {
            ans = append(ans, []int{0, j})
        }
        if g[m-1][j] == '.' && bfs(m - 1, j, 2) {
            ans = append(ans, []int{m - 1, j})
        }
    }

    for i := 1; i < m - 1; i++ {
        if g[i][0] == '.' && bfs(i, 0, 3) {
            ans = append(ans, []int{i, 0})
        }
        if g[i][n-1] == '.' && bfs(i, n - 1, 1) {
            ans = append(ans, []int{i, n - 1})
        }
    }
    
    return
}