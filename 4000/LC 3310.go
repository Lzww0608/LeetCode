func remainingMethods(n int, k int, invocations [][]int) []int {
    g := make([][]int, n)
    fa := make([]int, n)
    for i := 0; i < n; i++ {
        fa[i] = i
    }
    
    var find func(int) int
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            fa[rx] = ry
        }
    }
    for _, e := range invocations {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        merge(a, b)
    }

    vis := make([]bool, n)
    var dfs func(u int) 
    dfs = func(u int) {
        for _, v := range g[u] {
            if !vis[v] {
                vis[v] = true
                dfs(v)
            }
        } 
    }
    vis[k] = true
    dfs(k)

    res := make([]int, 0, n)
    for i := 0; i < n; i++ {
        if !vis[i] {
            if find(k) == find(i) {
                ans := make([]int, n)
                for i := 0; i < n; i++ {
                    ans[i] = i
                }
                return ans
            }
            res = append(res, i)
        }
    }

    return res
}