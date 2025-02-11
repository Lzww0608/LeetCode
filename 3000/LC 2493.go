func magnificentSets(n int, edges [][]int) (ans int) {
    g := make([][]int, n)
    for _, edge := range edges {
        a, b := edge[0] - 1, edge[1] - 1
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    bfs := func(start int) (d int) {
        q := []int{start}
        vis := make([]bool, n)
        vis[start] = true
        for len(q) > 0 {
            d++
            for sz := len(q); sz > 0; sz-- {
                u := q[0]
                q = q[1:]
                for _, v := range g[u] {
                    if !vis[v] {
                        vis[v] = true
                        q = append(q, v)
                    }
                }
            }
        }
        return
    }

    col := make([]int, n)
    nodes := []int{}
    var dfs func(int, int) bool
    dfs = func(u, c int) bool {
        if col[u] != 0 {
            return col[u] == c + 1
        }
        nodes = append(nodes, u)
        col[u] = c + 1
        for _, v := range g[u] {
            if !dfs(v, c ^ 1) {
                return false
            }
        }
        return true
    }

    for i := 0; i < n; i++ {
        if col[i] != 0 {
            continue
        }
        nodes = nil
        if !dfs(i, 0) {
            return -1
        }
        cur := 0
        for _, x := range nodes {
            cur = max(cur, bfs(x))
        }
        ans += cur
    }

    return
}