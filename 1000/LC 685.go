func findRedundantDirectedConnection(edges [][]int) []int {
    n := len(edges)
    fa := make([]int, n + 1)
    anc := make([]int, n + 1)
    for i := range fa {
        anc[i] = i
        fa[i] = i
    }
    
    var find func(int) int
    find = func(x int) int {
        if x != anc[x] {
            anc[x] = find(anc[x])
        }
        return anc[x]
    }

    merge := func(x, y int) {
        anc[find(x)] = find(y)
    }

    conf, cyc := -1, -1
    for i, e := range edges {
        a, b := e[0], e[1]
        if fa[b] != b {
            conf = i 
        } else {
            fa[b] = a
            if find(a) == find(b) {
                cyc = i
            } else {
                merge(a, b)
            }
        }
    }

    if conf < 0 {
        return []int{edges[cyc][0], edges[cyc][1]}
    } else {
        e := edges[conf]
        if cyc >= 0 {
            return []int{fa[e[1]], e[1]}
        } else {
            return []int{e[0], e[1]}
        }
    }

    return nil
}
