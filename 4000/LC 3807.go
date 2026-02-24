func minCost(n int, edges [][]int, k int) int {
    g := make([][]pair, n)
    for _, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        g[u] = append(g[u], pair{v, w})
        g[v] = append(g[v], pair{u, w})
    }

    vis := make([]int, n)
    check := func(mid int) bool {
        q := []int{0}
        for t := k; len(q) > 0 && t > 0; t-- {
            for sz := len(q); sz > 0; sz-- {
                u := q[0]
                q = q[1:]
                for _, v := range g[u] {
                    if v.w > mid || vis[v.u] == mid {
                        continue
                    }

                    if v.u == n - 1 {
                        return true
                    }
                    vis[v.u] = mid
                    q = append(q, v.u)
                }
            }
        }
        return false
    }


    l, r := 0, int(1e9 + 1)
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
    u, w int
}

