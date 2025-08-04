func minimumFuelCost(roads [][]int, seats int) int64 {
    n := len(roads) + 1
    g := make([][]int, n)
    for _, road := range roads {
        u, v := road[0], road[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }

    ans := 0
    var dfs func(int, int) int 
    dfs = func(u, fa int) int {
        sz := 1
        for _, v := range g[u] {
            if v != fa {
                sz += dfs(v, u)
            }
        }
        if u == 0 {
            return sz
        }
        ans += (sz + seats - 1) / seats
        return sz
    }
    dfs(0, -1)

    return int64(ans)
}