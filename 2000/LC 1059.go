func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
    g := make([][]int, n)
    in := make([]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        g[v] = append(g[v], u)
        in[u]++
    }

    if in[destination] > 0 {
        return false
    }

    q := []int{destination}
    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        if u == source {
            return true
        }
        for _, v := range g[u] {
            if in[v]--; in[v] == 0 {
                q = append(q, v)
                
            }
        }
    }

    return false
}