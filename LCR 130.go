func wardrobeFinishing(m int, n int, cnt int) (ans int) {
    vis := map[int]bool{}
    count := func(x int) int {
        res := 0
        for x > 0 {
            res += x % 10
            x /= 10
        }
        return res
    }

    var dfs func(int, int) 
    dfs = func(i, j int) {
        if i >= 0 && i < m && j >= 0 && j < n && !vis[i * n + j] && count(i) + count(j) <= cnt {
            ans++
            vis[i * n + j] = true
            dfs(i + 1, j)
            dfs(i - 1, j)
            dfs(i, j + 1)
            dfs(i, j - 1)
        }
    }
    dfs(0, 0)

    return 
}