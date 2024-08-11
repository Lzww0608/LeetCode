func find(fa []int, x int) int {
    if x != fa[x] {
        fa[x] = find(fa, fa[x])
    } 
    return fa[x]
}

func MST(n int, edges [][]int, f int) (res int) {
    fa := make([]int, n)
    for i := range fa {
        fa[i] = i
    }
    if f != -1 {
        res += edges[f][2]
        u := find(fa, edges[f][0])
        v := find(fa, edges[f][1])
        fa[u] = v
    }

    q := make([]int, len(edges))
    for i := range q {
        q[i] = i
    }
    sort.Slice(q, func(i, j int) bool {
        return edges[q[i]][2] < edges[q[j]][2]
    })

    for _, x := range q {
        u := find(fa, edges[x][0])
        v := find(fa, edges[x][1])
        if u != v {
            res += edges[x][2]
        }
        fa[u] = v
    }

    return 
}


func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
    keys := []int{}
    fkeys := []int{}
    mst := MST(n, edges, -1)
    for i := range edges {
        ok := false
        edges[i][2]++
        if MST(n, edges, -1) > mst {
            ok = true
        }
        edges[i][2]--
        if ok {
            keys = append(keys, i)
        }
        if !ok && MST(n, edges, i) == mst {
            fkeys = append(fkeys, i)
        }
    }

    return [][]int{keys, fkeys}
}