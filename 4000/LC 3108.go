func minimumCost(n int, edges [][]int, query [][]int) []int {
    fa := make([]int, n)
    sz := make([]int, n)
    for i := range fa {
        fa[i] = i
        sz[i] = -1
    }

    var find func(int) int 
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }

    merge := func(x, y, w int) {
        rx, ry := find(x), find(y)
        fa[rx] = ry
        sz[ry] &= sz[rx] & w
    }
    for _, edge := range edges {
        merge(edge[0], edge[1], edge[2])
    }

    ans := make([]int, len(query))
    for i, q := range query {
        a, b := find(q[0]), find(q[1])
        if a != b {
            ans[i] = -1
        } else {
            ans[i] = sz[a]
        }
    }

    return ans
}