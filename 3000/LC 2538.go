func maxOutput(n int, edges [][]int, price []int) int64 {
    g := make([][]int, n)
    for _, v := range edges {
        a, b := v[0], v[1]
        g[a] = append(g[a], b);
        g[b] = append(g[b], a)
    }

    ans := 0
    var dfs func(int, int) (int, int)
    dfs = func(u, fa int) (s1, s2 int) {
        x := price[u]
        s2 = x 
        for _, v := range g[u] {
            if v == fa {
                continue
            }
            a, b := dfs(v, u)
            ans = max(ans, s1 + b, s2 + a)
            s1 = max(s1, a + x)
            s2 = max(s2, b + x)
        }

        return s1, s2
    }

    dfs(0, -1)
    return int64(ans)
}