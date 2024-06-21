func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    fa, rk := make([]int, n + 1), make([]int, n + 1)
    cnt := n
    for i := range fa {
        fa[i] = i
        rk[i] = 1
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
            cnt--
            if rk[ry] <= rk[rx] {
                fa[ry] = rx
            } else {
                fa[rx] = ry;
            }
            if rk[rx] == rk[ry] {
                rk[rx]++
            }
        }
    }

    for i := range isConnected {
        for j := 0; j < i; j++ {
            x := isConnected[i][j]
            if x == 1 && find(i + 1) != find(j + 1) {
                merge(i + 1, j + 1)
            }
        }
    }

    return cnt
}
