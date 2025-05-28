func build(edge [][]int) [][]int {
    n := len(edge)
    g := make([][]int, n + 1)
    for _, v := range edge {
        a, b := v[0], v[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    return g
}

func dfs(u, fa, k int, g [][]int) (cnt int) {
    if k < 0 {
        return
    }

    cnt = 1
    for _, v := range g[u] {
        if v != fa {
            cnt += dfs(v, u, k - 1, g)
        }
    }

    return 
}

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
    mx := 0
    if k > 0 {
        g := build(edges2)
        m := len(edges2) + 1
        for i := 0; i < m; i++ {
            mx = max(mx, dfs(i, -1, k - 1, g))
        }
    }

    n := len(edges1) + 1
    ans := make([]int, n)
    g := build(edges1)
    for i := 0; i < n; i++ {
        ans[i] = dfs(i, -1, k, g) + mx
    }

    return ans
}