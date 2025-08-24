func minRunesToAdd(n int, crystals []int, flowFrom []int, flowTo []int) (ans int) {
    g := make([][]int, n)
    ug := make([][]int, n)
    in := make([]int, n)
    vis := make([]int, n)
    for i := range flowFrom {
        g[flowFrom[i]] = append(g[flowFrom[i]], flowTo[i])
        ug[flowFrom[i]] = append(ug[flowFrom[i]], flowTo[i])
        ug[flowTo[i]] = append(ug[flowTo[i]], flowFrom[i])
        in[flowTo[i]]++
    }

    var dfs func(int, bool) 
    dfs = func(u int, rev bool) {
        if vis[u] != 0 {
            return
        }
        vis[u] = 1
        if rev {
            for _, v := range ug[u] {
                dfs(v, rev)
            }
        } else {
            for _, v := range g[u] {
                dfs(v, rev)
            }
        }
    }

    for _, x := range crystals {
        dfs(x, false)
    }

    for i, x := range in {
        if x == 0 && vis[i] == 0 {
            dfs(i, false)
            ans++
        }
    }

    for i := 0; i < n; i++ {
        if vis[i] == 0 {
            dfs(i, true)
            ans++
        }
    }

    return
}