func minTime(n int, edges [][]int, hasApple []bool) (ans int) {
    g := make([][]int, n)
    for _, e := range edges {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    var dfs func(int, int) bool 
    dfs = func(idx, fa int) (res bool) {
        if len(g[idx]) == 1 && g[idx][0] == fa {
            return hasApple[idx]
        }
        for _, x := range g[idx] {
            if x != fa && dfs(x, idx) {
                res = true
                ans += 2
            }
        }

        return hasApple[idx] || res
    }
    dfs(0, -1)

    return
}