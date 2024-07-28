func calcEquation(equations [][]string, values []float64, queries [][]string) (ans []float64) {
    n := len(equations)
    m := map[string]int{}
    f := make([]int, n << 1)
    w := make([]float64, n << 1)
    t := 0

    for _, e := range equations {
        if _, ok := m[e[0]]; !ok {
            m[e[0]] = t
            t++
        }
        if _, ok := m[e[1]]; !ok {
            m[e[1]] = t
            t++
        }
    }

    for i := 0; i < t; i++ {
        f[i] = i
        w[i] = 1.0
    }

    var find func(int) int
    find = func(x int) int {
        if x != f[x] {
            tmp := f[x]
            f[x] = find(f[x])
            w[x] *= w[tmp]
        }
        return f[x]
    }

    merge := func(x, y int, val float64) {
        rx, ry := find(x), find(y)
        if rx != ry {
            f[rx] = ry
            w[rx] = val * w[y] / w[x]
        }
    }

    for i, e := range equations {
        a, b := m[e[0]], m[e[1]]
        merge(a, b, values[i])
    }

    for _, q := range queries {
        _, ok1 := m[q[0]]
        _, ok2 := m[q[1]]
        if !ok1 || !ok2 {
            ans = append(ans, -1)
            continue
        }
        a, b := find(m[q[0]]), find(m[q[1]])
        if a != b {
            ans = append(ans, -1)
        } else {
            ans = append(ans, w[m[q[0]]] / w[m[q[1]]])
        }
    }

    return 
}