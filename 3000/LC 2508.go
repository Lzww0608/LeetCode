func isPossible(n int, edges [][]int) bool {
    g := make([]map[int]struct{}, n)
    for i := 0; i < n; i++ {
        g[i] = map[int]struct{}{}
    }
    deg := make([]int, n)
    for _, edge := range edges {
        u, v := edge[0] - 1, edge[1] - 1
        g[u][v] = struct{}{}
        g[v][u] = struct{}{}
        deg[u]++
        deg[v]++
    }

    a := []int{}
    for i, x := range deg {
        if x & 1 == 1 {
            a = append(a, i)
        }
    }

    if len(a) & 1 == 1 || len(a) > 4 {
        return false
    } else if len(a) == 0 {
        return true
    }

    if len(a) == 2 {
        u, v := a[0], a[1]
        _, ok := g[u][v]
        if !ok {
            return true
        }
        for i := 0; i < n; i++ {
            if i != u && i != v {
                _, ok1 := g[u][i]
                _, ok2 := g[v][i]
                if !ok1 && !ok2 {
                    return true
                }
            }
        }
        return false
    } else {
        a, b, c, d := a[0], a[1], a[2], a[3]
        if _, ok1 := g[a][b]; !ok1 {
            if _, ok2 := g[c][d]; !ok2 {
                return true
            }
        } 
        if _, ok1 := g[a][c]; !ok1 {
            if _, ok2 := g[b][d]; !ok2 {
                return true
            }
        } 
        if _, ok1 := g[a][d]; !ok1 {
            if _, ok2 := g[c][b]; !ok2 {
                return true
            }
        } 
        return false
    }
}