func maximumScore(scores []int, edges [][]int) int {
    n := len(scores)
    ans := -1
    g := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }

    for i := 0; i < n; i++ {
        if len(g[i]) < 3 {
            continue
        }
        sort.Slice(g[i], func(p, q int) bool {
            return scores[g[i][p]] > scores[g[i][q]]
        })
        g[i] = g[i][:3]
    }

    for _, edge := range edges {
        u, v := edge[0], edge[1]
        for _, x := range g[u] {
            if x == v {
                continue
            }
            for _, y := range g[v] {
                if y != u && y != x {
                    ans = max(ans, scores[x] + scores[y] + scores[u] + scores[v])
                }
            }
        }
    }

    return ans
}