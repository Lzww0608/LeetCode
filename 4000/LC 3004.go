func maximumSubtreeSize(edges [][]int, colors []int) (ans int) {
    n := len(colors)
    g := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }

    var dfs func(int, int) int 
    dfs = func(u, fa int) (res int) {
        res = 1
        for _, v := range g[u] {
            if v == fa {
                continue
            }
            t := dfs(v, u)
            if res != -1 {
                if colors[u] != colors[v] {
                    res = -1
                } else {
                    res += t 
                }
            }
        }
        ans = max(ans, res)
        return
    }
    dfs(0, -1)
    return
}