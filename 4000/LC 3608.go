func minTime(n int, edges [][]int, k int) int {
    g := make([][]pair, n)
    for _, edge := range edges {
        a, b, c := edge[0], edge[1], edge[2]
        g[a] = append(g[a], pair{b, c})
        g[b] = append(g[b], pair{a, c})
    }

    check := func(mid int) bool {
        vis := make([]bool, n)
        cnt := 0

        var dfs func(int, int) 
        dfs = func(u, fa int) {
            if vis[u] {
                return
            }
            vis[u] = true
            for _, v := range g[u] {
                if v.x != fa && v.t > mid {
                    dfs(v.x, u)
                }
            }
        }
        for i := range n {
            if vis[i] {
                continue
            }
            cnt++
            dfs(i, -1)
        }

        return cnt >= k
    }
    
    l, r := -1, int(1e9) + 1
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid
        }
    }

    return r
}

type pair struct {
    x, t int
}