func validTree(n int, edges [][]int) bool {
    if len(edges) != n - 1 {
        return false
    }
    fa := make([]int, n)
    for i := range fa {
        fa[i] = i
    }

    var find func(int) int
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }

        return fa[x]
    }

    merge := func(x, y int) bool {
        rx, ry := find(x), find(y)
        if rx != ry {
            fa[rx] = ry
            return true
        } else {
            return false
        }
    }

    for _, e := range edges {
        u, v := e[0], e[1]
        if !merge(u, v) {
            return false
        }
    }

    return true
}