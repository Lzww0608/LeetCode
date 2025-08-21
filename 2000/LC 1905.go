func countSubIslands(a [][]int, b [][]int) (ans int) {
    m, n := len(a), len(a[0])
    vis := make([]bool, m * n)

    var dfs func(i, j int, valid bool) bool 
    dfs = func(i, j int, valid bool) (res bool) {
        if i < 0 || i >= m || j < 0 || j >= n || vis[i * n + j] || b[i][j] != 1 {
            return true
        }
        vis[i * n + j] = true

        res = true
        if valid {
            if a[i][j] != b[i][j] {
                valid = false
                res = false
            }
        } else {
            res = false
        }
        res = dfs(i + 1, j, valid) && res
        res = dfs(i - 1, j, valid) && res
        res = dfs(i, j + 1, valid) && res
        res = dfs(i, j - 1, valid) && res

        return
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if b[i][j] == 1 && !vis[i * n + j] {
                if dfs(i, j, true) {
                    ans++
                }
            }
        }
    }

    return
}