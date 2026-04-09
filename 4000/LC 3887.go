func numberOfEdgesAdded(n int, edges [][]int) (ans int) {
    fa := make([]int, n)
    dis := make([]int, n)
    for i := range fa {
        fa[i] = i
    }

    var find func(int) int 
    find = func(x int) int {
        if fa[x] != x {
            root := find(fa[x])
            dis[x] ^= dis[fa[x]]
            fa[x] = root
        }

        return fa[x]
    }

    merge := func(x, y, val int) bool {
        rx, ry := find(x), find(y)
        if rx == ry {
            return dis[x] ^ dis[y] == val
        }
        dis[rx] = val ^ dis[x] ^ dis[y]
        fa[rx] = ry 
        return true
    }

    for _, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        if merge(u, v, w) {
            ans++
        }
    }

    return
}