func findCircleNum(isConnected [][]int) (ans int) {
    n := len(isConnected)

    pa := make([]int, n)
    for i := range pa {
        pa[i] = i
    }

    var find func(x int) int
    find = func(x int) int {
        if x != pa[x] {
            pa[x] = find(pa[x])
        }
        return pa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            pa[ry] = rx
        }
    }

    for i := range isConnected {
        for j, x := range isConnected[i] {
            if x == 1 {
                merge(i, j)
            }
        }
    }


    for i, x := range pa {
        if x == i {
            ans++
        }
    }

    return 
}