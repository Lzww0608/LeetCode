func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    f := make([]int, n)
    r := make([]int, n)
    for i := range f {
        f[i] = i
    }

    var find func(int) int 
    find = func(x int) int {
        if f[x] != x {
            f[x] = find(f[x])
        }
        return f[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            if r[rx] < r[ry] {
                rx, ry = ry, rx
            }
            f[ry] = rx
            if r[rx] != r[ry] {
                r[rx]++
            }
        }
    }

    for _, e := range edges {
        a, b := e[0] - 1, e[1] - 1
        if find(a) == find(b) {
            return []int{a + 1, b + 1}
        }
        merge(a, b)
    }

    return nil
}