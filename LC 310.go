func findMinHeightTrees(n int, edges [][]int) (ans []int) {
    ans = append(ans, 0)
    g := make([][]int, n)
    deg := make([]int, n)

    for _, e := range edges {
        a,  b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
        deg[a]++
        deg[b]++
    }

    q := []int{}
    for i, x := range deg {
        if x == 1 {
            deg[i]--
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        ans = append([]int(nil), q...)
        for i := len(q); i > 0; i-- {
            t := q[0]
            q = q[1:]
            for _, x := range g[t] {
                deg[x]--
                if deg[x] == 1 {
                    q = append(q, x)
                }
            }
        }
    }

    return
}
