func maxScore(n int, edges [][]int) int64 {
    g := make([][]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        g[u] = append(g[u], v)
        g[v] = append(g[v], u)
    }

    a := make([]int, n)
    q := make([]int, 0, n)
    ans := 0
    if len(edges) == n {
        a[0] = n
        cur := n - 1
        q = append(q, 0)
        for len(q) > 0 {
            i := q[0]
            q = q[1:]
            for _, j := range g[i] {
                if a[j] != 0 {
                    continue
                }
                ans += a[i] * cur 
                a[j] = cur 
                cur--
                q = append(q, j)
            }
        }
        ans += 2
        return int64(ans)
    } 

    cur := n
    for i := range g {
        if len(g[i]) == 1 {
            q = append(q, i)
            a[i] = cur 
            cur--
        }
    }

    for len(q) > 0 {
        i := q[0]
        q = q[1:]
        for _, j := range g[i] {
            if a[j] != 0 {
                continue
            }
            ans += a[i] * cur 
            a[j] = cur
            cur--
            q = append(q, j)
        }
    }

    ans += n * (n - 1)
    return int64(ans)
}