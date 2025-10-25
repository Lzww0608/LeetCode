func minCost(n int, edges [][]int, k int) (ans int) {
    if k == n {
        return 0
    }

    sort.Slice(edges, func(i, j int) bool {
        return edges[i][2] < edges[j][2]
    })

    fa := make([]int, n)
    cnt := n
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

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            cnt--
            fa[rx] = ry
        }
    }

    pre := -1
    for _, v := range edges {
        merge(v[0], v[1])
        if cnt < k {
            break
        } else if pre != cnt {
            ans = v[2]
            pre = cnt
        }
    }

    return
}