func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
    forbid := make([][]bool, n)
    for i := range forbid {
        forbid[i] = make([]bool, n)
    }
    for _, v := range restrictions {
        a, b := v[0], v[1]
        forbid[a][b] = true
        forbid[b][a] = true
    }

    fa := make([]int, n)
    sz := make([]int, n)
    for i := range fa {
        fa[i] = i
        sz[i] = 1
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
        if rx == ry {
            return true
        }

        if forbid[rx][ry] || forbid[ry][rx] {
            return false
        }

        if sz[rx] < sz[ry] {
            rx, ry = ry, rx
        }

        fa[ry] = rx
        for k := 0; k < n; k++ {
            if forbid[ry][k] {
                forbid[rx][k], forbid[k][rx] =  true, true
            }
            if forbid[rx][k] {
                forbid[ry][k], forbid[k][ry] =  true, true
            }
        }

        return true
    }

    ans := make([]bool, len(requests))
    for i, v := range requests {
        x, y := find(v[0]), find(v[1])
        if merge(x, y) {
            ans[i] = true
        }
    }

    return ans
}