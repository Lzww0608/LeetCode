func longestLine(a [][]int) (ans int) {
    m, n := len(a), len(a[0])
    
    bfs := func(i, j, x, y int) (cnt int) {
        for i >= 0 && i < m && j >= 0 && j < n && a[i][j] == 1 {
            cnt++
            i += x
            j += y 
        }
        return
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if a[i][j] == 1 {
                ans = max(ans, bfs(i, j, 0, 1))
                ans = max(ans, bfs(i, j, 1, 0))
                ans = max(ans, bfs(i, j, 1, 1))
                ans = max(ans, bfs(i, j, 1, -1))
            }
        }
    }

    return ans
}