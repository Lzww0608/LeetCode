const N int = 14
func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
    type edge struct {
        to, wt int
    }

    type pair struct {
        p   int
        cnt [26]int
    }

    g := make([][]edge, n)
    for _, e := range edges {
        u, v, w := e[0], e[1], e[2] - 1
        g[u] = append(g[u], edge{v, w})
        g[v] = append(g[v], edge{u, w})
    }

    fa := make([][N]pair, n)
    depth := make([]int, n)
    
    var dfs func(int, int, int)
    dfs = func(v, p, d int) {
        fa[v][0].p = p
        depth[v] = d
        for _, e := range g[v] {
            if u := e.to; u != p {
                fa[u][0].cnt[e.wt] = 1
                dfs(u, v, d + 1)
            }
        }
    }
    dfs(0, - 1, 0)
    
    // 倍增
    for i := 0; i < N - 1; i++ {
        for x := 0; x < n; x++ {
            if p := fa[x][i]; p.p != -1 {
                fa[x][i+1] = fa[p.p][i]
                for j := 0; j < 26; j++ {
                    fa[x][i+1].cnt[j] = p.cnt[j] + fa[p.p][i].cnt[j]
                }
            } else {
                fa[x][i+1].p = -1
            }
        }
    }

    f := func(u, v int) int {
        tot := depth[v] + depth[u]
        cnt := [26]int{}
        for depth[u] > depth[v] {
            u, v = v, u
        }

        for i := 0; i < N; i++ {
            if ((depth[v] - depth[u]) >> i) & 1 != 0 {
                p := fa[v][i]
                for j := 0; j < 26; j++ {
                    cnt[j] += p.cnt[j]
                }
                v = p.p
            } 
        }

        if v != u {
            for i := N - 1; i >= 0; i-- {
                if pu, pv := fa[u][i], fa[v][i]; pu.p != pv.p {
                    for j := 0; j < 26; j++ {
                        cnt[j] += pv.cnt[j] + pu.cnt[j]
                    }
                    u, v = pu.p, pv.p
                }
            }

            for j := 0; j < 26; j++ {
                cnt[j] += fa[u][0].cnt[j] + fa[v][0].cnt[j]
            }

            u = fa[u][0].p
        }

        return tot - depth[u] * 2 - slices.Max(cnt[:])
    }

    ans := make([]int, len(queries))
    for i, q := range queries {
        ans[i] = f(q[0], q[1])
    }
    return ans
}