func numTimesAllBlue(flips []int) (ans int) {
    n := len(flips)
    fa := make([]int, n + 1)
    sz := make([]int, n + 1)
    for i := range fa {
        fa[i] = i
        sz[i] = 1
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
            if rx < ry {
                rx, ry = ry, rx
            }
            fa[rx] = ry
            sz[ry] += sz[rx]
        }
    }

    for i, x := range flips {
        merge(x, x - 1)
        y := find(x)
        if y == 0 && sz[y] == i + 2 {
            ans++
        }
    }

    return
}