func eventualSafeNodes(graph [][]int) []int {
    n := len(graph)
    ans := make([]int, 0, n)
    in := make([]int, n)
    g := make([][]int, n)
    q := []int{}

    for i, v := range graph {
        if len(v) == 0 {
            ans = append(ans, i)
            q = append(q, i)
        }

        for _, j := range v {
            in[i]++
            g[j] = append(g[j], i)
        }
    }

    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        for _, v := range g[u] {
            in[v]--
            if in[v] == 0 {
                q = append(q, v)
                ans = append(ans, v)
            }
        }
    }

    sort.Ints(ans)
    return ans
}