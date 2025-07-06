func minEdgeReversals(n int, edges [][]int) []int {
    type pair struct {
        node, val int 
    }
    g := make([][]pair, n)
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        g[a] = append(g[a], pair{b, 1})
        g[b] = append(g[b], pair{a, -1})
    }

    ans := make([]int, n)
    var dfs func(int, int)
    dfs = func(u, fa int) {
        for _, v := range g[u] {
            if v.node != fa {
                if v.val < 0 {
                    ans[0]++
                }
                dfs(v.node, u)
            }
        }
    }
    dfs(0, -1)

    var dfs1 func(int, int)
    dfs1 = func(u, fa int) {
        for _, v := range g[u] {
            if v.node != fa {
                ans[v.node] = ans[u] + v.val
                dfs1(v.node, u)
            }
        }
    }
    dfs1(0, -1)

    return ans
}