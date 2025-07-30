func timeTaken(edges [][]int) []int {
    n := len(edges)
    g := make([][]int, n + 1)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }

    ans := make([]int, n + 1)
    f := make([]struct{mx, sec, p int}, n + 1)
    var dfs func(int, int) int 
    dfs = func(u, fa int) int {
        for _, v := range g[u] {
            if v != fa {
                p := dfs(v, u) + 2 - v % 2
                if p >= f[u].mx {
                    f[u].sec = f[u].mx 
                    f[u].mx = p 
                    f[u].p = v
                } else if p > f[u].sec {
                    f[u].sec = p 
                }
            }
        }

        return f[u].mx
    }
    dfs(0, -1)

    var reroot func(int, int, int)
    reroot = func(u, fa, mx int) {
        t := f[u]
        ans[u] = max(mx, t.mx)
        for _, v := range g[u] {
            if v != fa {
                w := 2 - u % 2 
                if v == t.p {
                    reroot(v, u, max(mx, t.sec) + w)
                } else {
                    reroot(v, u, max(mx, t.mx) + w)
                }
            }
        }
    }
    reroot(0, -1, 0)
    return ans
}