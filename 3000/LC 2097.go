func validArrangement(pairs [][]int) [][]int {
    g := make(map[int][]int)
    in := make(map[int]int)

    for _, p := range pairs {
        u, v := p[0], p[1]
        g[u] = append(g[u], v)
        in[v]++
    }

    start := 0
    for i, v := range g {
        if len(v) == in[i] + 1 {
            start = i 
            break
        }
        start = i 
    }

    m := len(pairs)
    ans := make([][]int, 0, m)

    var dfs func(int)
    dfs = func(u int) {
        for len(g[u]) > 0 {
            v := g[u][0]
            g[u] = g[u][1:]
            dfs(v)
            ans = append(ans, []int{u, v})
        }
    }
    dfs(start)

    for i := 0; i < m / 2; i++ {
        ans[i], ans[m-1-i] = ans[m-1-i], ans[i]
    }

    return ans
}