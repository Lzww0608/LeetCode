func minimumThreshold(n int, edges [][]int, source int, target int, k int) int {
    g := make([][]pair, n)
    for _, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        g[u] = append(g[u], pair{v, w})
        g[v] = append(g[v], pair{u, w})
    }
    
    m := make([]int, n)
    check := func(mid int) bool {
        for i := range m {
            m[i] = k + 1
        }
        m[source] = 0
        q := [][2]int{{source, 0}}
        for len(q) > 0 {
            cur := q[0]
            u, cnt := cur[0], cur[1]
            q = q[1:] 
            if cnt > m[u] {
                continue
            }
            for _, v := range g[u] {
                cur := cnt 
                if v.w > mid {
                    cur++
                }

                if cur < m[v.v] {
                    m[v.v] = cur 
                    q = append(q, [2]int{v.v, cur})
                }
            }
        }

        return m[target] <= k
    }
    
    l, r := -1, int(1e9 + 1)
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid 
        } else {
            l = mid 
        }
    }

    if r > int(1e9) {
        return -1
    }
    return r
}

type pair struct {
    v, w int 
}

