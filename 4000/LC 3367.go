func maximizeSumOfWeights(edges [][]int, k int) int64 {
    n := len(edges) + 1
    g := make([][][2]int, n + 1)
    sum := 0
    for _, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        g[u] = append(g[u], [2]int{v, w})
        g[v] = append(g[v], [2]int{u, w})
        sum += w
    }

    ret := true 
    for _, v := range g {
        if len(v) > k {
            ret = false
            break
        }
    }

    if ret {
        return int64(sum)
    }

    var dfs func(int, int) (int, int) 
    dfs = func(u, fa int) (notChoose, choose int) {
        a := []int{}
        for _, t := range g[u] {
            v, w := t[0], t[1]
            if v == fa {
                continue
            }
            nc, c := dfs(v, u)
            notChoose += nc 
            if d := w + c - nc; d > 0 {
                a = append(a, d)
            }
        }

        sort.Slice(a, func(i, j int) bool {
            return a[i] > a[j]
        })
        for i := range min(len(a), k - 1) {
            notChoose += a[i]
        }
        choose = notChoose
        if len(a) >= k {
            notChoose += a[k - 1]
        }
        return
    }

    nc, _ := dfs(0, -1)
    return int64(nc)
}