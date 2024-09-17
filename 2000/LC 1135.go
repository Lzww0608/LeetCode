func minimumCost(n int, connections [][]int) int {
    fa := make([]int, n)
    cost, cnt := 0, 0
    for i := range fa {
        fa[i] = i
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
            cnt++
        }
        return
    }

    sort.Slice(connections, func(i, j int) bool {
        return connections[i][2] < connections[j][2]
    })

    
    for _, e := range connections {
        if cnt == n - 1 {
            break
        }

        u, v, w := e[0] - 1, e[1] - 1, e[2]
        if find(u) == find(v) {
            continue
        }
        merge(u, v)
        cost += w
    }
    if cnt != n - 1 {
        return -1
    }

    return cost
}