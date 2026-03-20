func maxActivated(points [][]int) int {
    n := len(points)
    fa := make([]int, n)
    sz := make([]int, n)
    for i := range n {
        fa[i], sz[i] = i, 1
    }

    var find func(x int) int 
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx == ry {
            return 
        } else if sz[rx] < sz[ry] {
            rx, ry = ry, rx
        }

        sz[rx] += sz[ry]
        fa[ry] = rx
    }

    ans := 2
    xs := make(map[int]int)
    ys := make(map[int]int)
    a := []int{}
    for i, v := range points {
        x, y := v[0], v[1]
        if t, ok := xs[x]; ok {
            merge(i, t)
        } else {
            xs[x] = i
        }
        if t, ok := ys[y]; ok {
            merge(i, t)
        } else {
            ys[y] = i
        }
    }

    for i := range points {
        if i == find(i) {
            a = append(a, i)
        }
    }

    sort.Slice(a, func(i, j int) bool {
        return sz[a[i]] > sz[a[j]]
    })
    
    if len(a) > 1 {
        ans = sz[a[0]] + sz[a[1]] + 1
    } else {
        ans = sz[a[0]] + 1
    }

    return ans
}