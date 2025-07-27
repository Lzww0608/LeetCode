func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
    d1 := solve(edges1)
    d2 := solve(edges2)

    return max(d1, d2, (d1 + 1) / 2 + (d2 + 1) / 2 + 1)
}

func solve(edges [][]int) int {
    n := len(edges) + 1
    g := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }

    d := 0
    var dfs func(int, int) int 
    dfs = func(u, fa int) int {
        a, b := 0, 0 
        for _, v := range g[u] {
            if v != fa {
                x := dfs(v, u)
                if x > a {
                    a, b = x, a
                } else if x > b {
                    b = x
                }
            }
        }
        d = max(d, a + b)
        return max(a, b) + 1
    } 
    dfs(0, -1)
    
    return d
}