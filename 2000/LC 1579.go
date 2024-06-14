func maxNumEdgesToRemove(n int, edges [][]int) (ans int) {
    fa_Bob := make([]int, n + 1)
    fa_Alice := make([]int, n + 1)
    for i := range fa_Alice {
        fa_Alice[i] = i
        fa_Bob[i] = i
    }

    var find func([]int, int) int
    find = func(fa []int, x int) int {
        if x != fa[x] {
            fa[x] = find(fa, fa[x])
        }
        return fa[x]
    }

    merge := func(fa []int, x, y int) {
        rx, ry := find(fa, x), find(fa, y)
        if rx != ry {
            fa[rx] = ry
        }
    }

    a, b := 0, 0
/*    sort.Slice(edges, func(i, j int) bool {
        return edges[i][0] > edges[j][0]
    })
    */
    for _, e := range edges {
        types, u, v := e[0], e[1], e[2]
        if types == 3 {
            if find(fa_Alice, u) != find(fa_Alice, v) {
                merge(fa_Alice, u, v)
                merge(fa_Bob, u, v)
                a++
                b++
            } else {
                ans++
            }
        } 
    }
    for _, e := range edges {
        types, u, v := e[0], e[1], e[2]
        if types == 1 {
            if find(fa_Alice, u) != find(fa_Alice, v) {
                merge(fa_Alice, u, v)
                a++
            } else {
                ans++
            }
        } else if types == 2 {
            if find(fa_Bob, u) != find(fa_Bob, v) {
                merge(fa_Bob, u, v)
                b++
            } else {
                ans++
            }
        }
    }
    if a != n - 1 || b != n - 1 {
        return -1
    }

    return
}
