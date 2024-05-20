func countPairs(n int, edges [][]int) (ans int64) {
    fa := make([]int, n)
    sz := make([]int, n)
    for i := range fa {
        fa[i], sz[i] = i, 1
    }

    var find func(int) int
    find = func(x int) int {
        if x != fa[x] {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }
    
    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            fa[rx] = ry
            sz[ry] += sz[rx]
        }
    }

    for _, e := range edges {
        a, b := e[0], e[1]
        merge(a, b)
    }

    var sum int64 = 0
    for i := range fa {
        if i == fa[i] {
            ans += sum * int64(sz[i])
            sum += int64(sz[i])
        }
    }

    return
}